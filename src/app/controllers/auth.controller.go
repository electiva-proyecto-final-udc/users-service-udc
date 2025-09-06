package controllers

import (
	"encoding/json"
	"net/http"
	"user-service-ucd/src/app/models"
	"user-service-ucd/src/app/services"
	"user-service-ucd/src/common"
)

type AuthController struct {
	as *services.AuthService
}

func NewAuthController(as *services.AuthService) *AuthController {
	return &AuthController{as: as}
}

// login Login
// @Summary      Logearse como técnico o admin
// @Description  Retorna un token JWT si las credenciales son correctas
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      models.LoginRequest  true  "Credenciales de login"
// @Success      200 {object} common.ApiResponse{data=common.AuthResult} "OK"
// @Failure      400 {object} common.ApiResponse{error=common.ErrorResponse}
// @Failure      401 {object} common.ApiResponse{error=common.ErrorResponse}
// @Failure      500 {object} common.ApiResponse{error=common.ErrorResponse}
// @Router       /auth/login [post]
func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.JSONResponse(w, http.StatusBadRequest, common.ApiResponse{
			Error: &common.ErrorResponse{
                Code: 400,
				Message: "Invalid request body",
				Details: err.Error(),
			},
		})
		return
	}

	response, err := ac.as.Login(req.Username, req.Password)
	if err != nil {
		common.JSONResponse(w, http.StatusUnauthorized, common.ApiResponse{
			Error: &common.ErrorResponse{
                Code: 401,
				Message: "INVALID_CREDENTIALS",
			},
		})
		return
	}
    
	common.JSONResponse(w, http.StatusOK, common.ApiResponse{
		Message: "LOGIN_SUCCESSFUL",
		Data:    response,
	})
}
