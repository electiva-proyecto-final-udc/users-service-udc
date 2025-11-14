package dto

type LoginRequest struct {
	Username string `json:"username" example:"carlos.r"`
	Password string `json:"password" example:"secret123"`
}

