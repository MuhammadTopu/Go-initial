package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"task_m/dto"

	"task_m/config"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	db *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) GetPublicProfile(c *gin.Context) {
	userID := c.GetString("user_id")

	var profile dto.UserPublicProfile

	err := h.db.QueryRow(`
		SELECT id, username, full_name, avatar_url, bio, created_at
		FROM users
		WHERE id = $1
	`, userID).Scan(
		&profile.ID,
		&profile.Username,
		&profile.FullName,
		&profile.AvatarURL,
		&profile.Bio,
		&profile.JoinedAt,
	)

	if err != nil {
		fmt.Println("DB error:", err)
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "User not found"})
		return
	}

	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Failed to load the config: %v\n", err)
	}

	var imageBaseURL = cfg.URL.BaseURL + "/uploads/avatars/"

	fmt.Printf("DB avatar_url: %s\n", *profile.AvatarURL)

	if *profile.AvatarURL != "" {
		avatar := imageBaseURL + *profile.AvatarURL
		profile.AvatarURL = &avatar
	}

	c.JSON(http.StatusOK, profile)
}
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := c.GetString("user_id")
	var req dto.UpdateProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	_, err := h.db.Exec(`
		UPDATE users 
		SET full_name = $1, bio = $2, avatar_url = $3, 
		    location_lat = $4, location_lng = $5, location_address = $6,
		    updated_at = CURRENT_TIMESTAMP
		WHERE id = $7
	`, req.FullName, req.Bio, req.AvatarURL, req.LocationLat, req.LocationLng,
		req.LocationAddress, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse("Profile updated successfully", nil))
}
func (h *UserHandler) UploadAvatar(c *gin.Context) {
	userID := c.GetString("user_id")

	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Avatar file is required"})
		return
	}

	// file size limit (5MB)
	if file.Size > 5*1024*1024 {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Avatar file size exceeds 5MB"})
		return
	}

	// validate extension
	allowedExt := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
		".avif": true,
	}

	fileExt := strings.ToLower(filepath.Ext(file.Filename))

	if !allowedExt[fileExt] {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid image format"})
		return
	}

	// generate unique filename
	fileName := uuid.New().String() + fileExt
	savePath := fmt.Sprintf("uploads/avatars/%s", fileName)

	// save file
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to upload avatar"})
		return
	}

	avatarURL := fmt.Sprintf("/%s", savePath)

	// get old avatar
	var existingAvatar string
	err = h.db.QueryRow(`
		SELECT avatar_url FROM users WHERE id = $1
	`, userID).Scan(&existingAvatar)

	avatarFileName := filepath.Base(savePath)

	// update database
	_, err = h.db.Exec(`
		UPDATE users
		SET avatar_url = $1, updated_at = CURRENT_TIMESTAMP
		WHERE id = $2
	`, avatarFileName, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to update avatar"})
		return
	}

	// delete old avatar
	if existingAvatar != "" {
		oldPath := existingAvatar[1:]
		os.Remove(oldPath)
	}

	c.JSON(http.StatusOK, dto.SuccessResponse(
		"Avatar uploaded successfully",
		gin.H{"avatar_url": avatarURL},
	))
}
