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

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
    var req models.LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        common.JSONResponse(w, http.StatusBadRequest, common.ApiResponse{
            Error: &common.ErrorResponse{
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
                Message: "INVALID_CREDENTIALS",
            },
        })
        return
    }

    common.JSONResponse(w, http.StatusOK, common.ApiResponse{
        Message: "LOGIN_SUCCESSFUL",
        Data: response,
    })
}
