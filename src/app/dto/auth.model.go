package dto

import "github.com/golang-jwt/jwt/v5"

type UserInfo struct {
	ID       string `json:"ID"`
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Role     string `json:"Role"`
}

type CustomClaims struct {
	UserData UserInfo
	jwt.RegisteredClaims
}
