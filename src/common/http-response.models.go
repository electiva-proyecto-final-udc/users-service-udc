package common

type ApiResponse struct {
	Message    string         `json:"message,omitempty"`
	Data       interface{}    `json:"data,omitempty"`
	Error      *ErrorResponse `json:"error,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Details interface{} `json:"details"`
}

type AuthResult struct {
    Role  string
    Token string
}