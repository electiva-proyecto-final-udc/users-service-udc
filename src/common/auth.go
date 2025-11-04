package common

import (
	"user-service-ucd/src/app/dto"
)

func GenerateToken(userInfo dto.UserInfo) (string, error) {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...", nil
}
