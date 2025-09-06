package common

type ApiResponse struct {
	Message string         `json:"message,omitempty"`
	Data    interface{}    `json:"data,omitempty"`
	Error   *ErrorResponse `json:"error,omitempty"`
}

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

type AuthResult struct {
    Role  string `json:"role" example:"TECHNICIAN"` // Rol del usuario autenticado
    Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."` // JWT generado
}