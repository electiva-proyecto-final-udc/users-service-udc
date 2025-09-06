package common

import "user-service-ucd/src/app/models"

func GenerateToken(userInfo models.UserInfo) (string, error) {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...", nil
}
