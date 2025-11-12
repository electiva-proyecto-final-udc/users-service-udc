package common

import (
	"fmt"
	"os"
	"time"
	"user-service-ucd/src/app/dto"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userInfo dto.UserInfo) (string, error) {
	claims := dto.CustomClaims{
		UserData: userInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: userInfo.ID,
			Issuer:    "user_service_udc",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(4 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func VerifyJWT(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("INVALID TOKEN")
}
