package models

type LoginRequest struct {
	Username string `json:"Username" example:"carlos.r"`
	Password string `json:"Password" example:"secret123"`
}

type UserInfo struct {
	ID       string `json:"ID"`
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Role     string `json:"Role"`
}
