package services

import (
	"user-service-ucd/src/app/dto"
	"user-service-ucd/src/app/repository"
	"user-service-ucd/src/common"
)

type AuthService struct {
	ur *repository.UserRepository
}

func NewAuthService(ur *repository.UserRepository) *AuthService {
	return &AuthService{
		ur: ur,
	}
}

func (as *AuthService) Login(loginReq dto.LoginRequest) (authResult common.AuthResult, err error) {
	var userDataInfo dto.UserInfo

	userLogged, errorUser := as.ur.FindUserByUsernamePassword(loginReq)
	if errorUser != nil {
		return common.AuthResult{}, errorUser
	}

	userDataInfo.ID = userLogged.ID
	userDataInfo.Email = userLogged.Email
	userDataInfo.Username = userLogged.Username
	userDataInfo.Role = userLogged.Role

	token, err := common.GenerateToken(userDataInfo)
	if err != nil {
		return common.AuthResult{}, err
	}
	authResult.Token = token
	authResult.Role = userLogged.Role

	return authResult, nil
}
