package models

type LoginRequest struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type UserInfo struct {
	ID       string `json:"ID"`
	Username string `json:"Username"`
	Email    string `json:"Emiail"`
	Role     string `json:"Role"`
}
