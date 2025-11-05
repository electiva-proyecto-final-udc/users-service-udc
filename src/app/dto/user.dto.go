package dto

type LoginRequest struct {
	Username string `json:"Username" example:"carlos.r"`
	Password string `json:"Password" example:"secret123"`
}

