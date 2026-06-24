package dto

type LoginRequest struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
    ExpiresAt int64  `json:"expires_at"`
}