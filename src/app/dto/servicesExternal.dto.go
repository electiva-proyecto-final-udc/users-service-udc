package dto

type SendWelcomeEmailRequest struct {
	To       string `json:"to"`
	Username string `json:"username"`
	Password string `json:"password"`
}
