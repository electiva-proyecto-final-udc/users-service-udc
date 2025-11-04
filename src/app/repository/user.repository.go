package repository

import (
	"fmt"
	"user-service-ucd/src/app/dto"
	"user-service-ucd/src/app/models"
	"user-service-ucd/src/database"
	"user-service-ucd/utils"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// Constructor: se crea una nueva instancia
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: database.DB,
	}
}

func (ur *UserRepository) FindUserByUsernamePassword(userInfo dto.LoginRequest) (userData models.UserDataView, err error) {
	var userProfile models.UserProfileEntity
	var userDataView models.UserDataView

	err = ur.db.Where("username = ? OR email = ?", userInfo.Username, userInfo.Username).Find(&userDataView).Error

	if err != nil {
		return models.UserDataView{}, err
	}

	if userDataView == (models.UserDataView{}) {
		return models.UserDataView{}, fmt.Errorf("USER NOT FOUND")
	}

	if !userDataView.IsActive {
		return models.UserDataView{}, fmt.Errorf("USER NOT ACTIVE")
	}

	err = ur.db.Where("id = ?", userDataView.UserId).Find(&userProfile).Error
	if err != nil {
		return models.UserDataView{}, err
	}

	if !utils.CheckPasswordHash(userInfo.Password, userProfile.Password) {
		return models.UserDataView{}, fmt.Errorf("INVALID CREDENTIALS")
	}

	return userDataView, nil
}
