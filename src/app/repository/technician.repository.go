package repository

import (
	"encoding/json"
	"fmt"
	"user-service-ucd/src/app/dto"
	"user-service-ucd/src/app/models"
	"user-service-ucd/src/database"

	"gorm.io/gorm"
)

// Se define la estructura del repositorio
type TechnicianRepository struct {
	dbFake []dto.TechnicianDTO
	db     *gorm.DB
}

// Constructor: se crea una nueva instancia
func NewTechnicianRepository(db *gorm.DB) *TechnicianRepository {
	return &TechnicianRepository{
		dbFake: []dto.TechnicianDTO{},
		db:     database.DB,
	}
}

// Crear un nuevo técnico
func (tr *TechnicianRepository) CreateNewTechnician(technician dto.CreateTechnicianDTO) error {
	var personProfileData models.PersonProfile
	var userProfileData models.UserProfileEntity
	technicianData, _ := json.Marshal(technician)

	if err := json.Unmarshal(technicianData, &personProfileData); err != nil {
		return err
	}

	if err := json.Unmarshal(technicianData, &userProfileData); err != nil {
		return err
	}

	// INICIA LA TRANSACCIÓN
	tx := tr.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(&personProfileData).Error; err != nil {
		tx.Rollback()
		return err
	}
	userProfileData.PersonProfileID = personProfileData.ID

	if err := tx.Create(&userProfileData).Error; err != nil {
		tx.Rollback()
		return err
	}

	// CIERRA LA TRANSACCIÓN (Hace rollback si falla)
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Obtener todos los técnicos
func (tr *TechnicianRepository) GetAllTechnicians() ([]models.UserDataView, error) {
	var users []models.UserDataView
	err := tr.db.Where("role_code = ?", "3").Find(&users).Error
	return users, err
}

// Obtener un técnico por ID
func (tr *TechnicianRepository) GetTechnicianById(id string) (models.UserDataView, error) {
	var user models.UserDataView
	err := tr.db.Where("id = ? AND role_code = 3", id).Find(&user).Error

	if err != nil {
		return models.UserDataView{}, err
	}

	if user == (models.UserDataView{}) {
		return models.UserDataView{}, fmt.Errorf("TECHNICIAN NOT FOUND")
	}

	return user, nil
}

// Actualizar técnico por ID
func (tr *TechnicianRepository) UpdateTechnician(id string, updated dto.UpdateTechnicianDTO) error {
	var technicianUpdated models.UpdateTechnician
	techinicianData, _ := json.Marshal(updated)
	if err := json.Unmarshal(techinicianData, &technicianUpdated); err != nil {
		return err
	}

	result := tr.db.Model(&models.UpdateTechnician{}).
		Where("id = ?", id).
		Updates(technicianUpdated)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("TECHNICIAN NOT FOUND")
	}

	return nil
}

// Eliminar técnico por ID
func (tr *TechnicianRepository) DeleteTechnician(id string) error {
	tx := tr.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	result := tx.Model(models.PersonProfile{}).
		Where("id = ?", id).
		Delete(nil)

	if result.Error != nil {
		tx.Rollback()
		return fmt.Errorf("ERROR DELETING TECHNICIAN: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("TECHNICIAN NOT FOUND")
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// Cambiar contraseña de un técnico por Username
func (tr *TechnicianRepository) ChangePassword(changePasswordRequest dto.ChangePasswordDTO) error {
	result := tr.db.Model(&models.UserProfileEntity{}).
		Where("person_profile_id = ?", changePasswordRequest.UserId).
		Updates(map[string]interface{}{
			"password": changePasswordRequest.NewPassword,
		})

	if result.RowsAffected == 0 {
		return fmt.Errorf("USER NOT FOUND")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// FindByUsername busca por Username o Email
func (tr *TechnicianRepository) FindByUsername(usernameOrEmail string) (*dto.TechnicianDTO, error) {
	for _, t := range tr.dbFake {
		if t.Username == usernameOrEmail || t.Email == usernameOrEmail {
			return &t, nil
		}
	}
	return nil, fmt.Errorf("technician not found")
}

// SET ACTIVE
