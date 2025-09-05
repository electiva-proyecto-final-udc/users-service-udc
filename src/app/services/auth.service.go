package services

import (
	"errors"
	"user-service-ucd/src/app/models"
	"user-service-ucd/src/app/repository"
	"user-service-ucd/src/common"
	"user-service-ucd/utils"
)

type AuthService struct {
	technicianRepo *repository.TechnicianRepository
	admin          models.Admin // ADMIN QUEMADO POR AHORA
}

func NewAuthService(tr *repository.TechnicianRepository) *AuthService {
	return &AuthService{
		technicianRepo: tr,
		admin: models.Admin{
			ID:           "1",
			DocumentType: "CC",
			Document:     "1234567",
			Name:         "Juancho",
			Surname:      "Ramirez",
			Email:        "juancho@j.com",
			PhoneNumber:  "121234567",
			Username:     "juancho123",
			Password:     "123",
			Permissions:  "",
		},
	}
}

func (as *AuthService) Login(username, password string) (common.AuthResult, error) {
	role := ""
	var userInfo models.UserInfo

	if (username == as.admin.Username || username == as.admin.Email) && (password == as.admin.Password) {
		role = "ADMIN"
		userInfo = models.UserInfo{
			ID:       as.admin.ID,
			Username: as.admin.Username,
			Email:    as.admin.Email,
			Role:     role,
		}
	} else {
		technician, err := as.technicianRepo.FindByUsername(username)
		if err != nil {
			return common.AuthResult{}, errors.New("invalid username or password")
		}

		if utils.CheckPasswordHash(password, technician.Password) {
			role = "TECHNICIAN"
			userInfo = models.UserInfo{
				ID:       technician.ID,
				Username: technician.Username,
				Email:    technician.Email,
				Role:     role,
			}
		}
	}

	if (userInfo == models.UserInfo{}) {
		return common.AuthResult{}, errors.New("invalid username or password")
	}

	token, err := common.GenerateToken(userInfo)
	if err != nil {
		return common.AuthResult{}, err
	}
	return common.AuthResult{Role: role, Token: token}, nil
}
