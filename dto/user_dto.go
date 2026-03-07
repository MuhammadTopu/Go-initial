package dto

type UpdateProfileRequest struct {
	FullName        string   `json:"full_name"`
	Bio             string   `json:"bio"`
	AvatarURL       string   `json:"avatar_url"`
	LocationLat     *float64 `json:"location_lat"`
	LocationLng     *float64 `json:"location_lng"`
	LocationAddress string   `json:"location_address"`
}

type UserPublicProfile struct {
	ID        string  `json:"id"`
	Username  string  `json:"username"`
	FullName  string  `json:"full_name"`
	AvatarURL *string `json:"avatar_url"`
	Bio       *string `json:"bio"`
	JoinedAt  string  `json:"joined_at"`
}
