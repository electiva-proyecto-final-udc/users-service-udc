package common

type ApiResponse struct {
	Success    bool           `json:"success"`
	StatusCode int            `json:"statusCode"`
	Message    string         `json:"message,omitempty"`
	Data       interface{}    `json:"data,omitempty"`
	Error      *ErrorResponse `json:"error,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Details string `json:"details"`
}
