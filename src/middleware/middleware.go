package middleware

import (
	"context"
	"net/http"
	"strings"
	"user-service-ucd/src/common"
)

type ctxUserData struct{}

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			common.JSONResponse(w, http.StatusUnauthorized, common.ApiResponse{
				Error: &common.ErrorResponse{
					Code:    401,
					Message: "MISSING AUTHORIZATION HEADER",
				},
			})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			common.JSONResponse(w, http.StatusUnauthorized, common.ApiResponse{
				Error: &common.ErrorResponse{
					Code:    401,
					Message: "INVALID AUTHORIZATION HEADER",
				},
			})
			return
		}

		tokenStr := parts[1]

		// 👇 Aquí reutilizas tu método
		claims, err := common.VerifyJWT(tokenStr)
		if err != nil {
			common.JSONResponse(w, http.StatusUnauthorized, common.ApiResponse{
				Error: &common.ErrorResponse{
					Code:    401,
					Message: "INVALID OR EXPIRED TOKEN",
				},
			})
			return
		}

		// Guardar los claims en el contexto
		ctx := context.WithValue(r.Context(), ctxUserData{}, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Helper para obtener la data del usuario en los handlers
func GetUserData(r *http.Request) string {
	if data, ok := r.Context().Value(ctxUserData{}).(string); ok {
		return data
	}
	return ""
}
