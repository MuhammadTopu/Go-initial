package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"task_m/dto"

	"github.com/gin-gonic/gin"
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
