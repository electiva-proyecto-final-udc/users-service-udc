package repository

import (
	"fmt"
	"user-service-ucd/src/app/models"
	"user-service-ucd/src/database"

	"gorm.io/gorm"
)

// Se define la estructura del repositorio
type TechnicianRepository struct {
	dbFake []models.Technician
	db     *gorm.DB
}

// Constructor: se crea una nueva instancia
func NewTechnicianRepository(db *gorm.DB) *TechnicianRepository {
	return &TechnicianRepository{
		dbFake: []models.Technician{},
		db: database.DB,
	}
}

// Crear un nuevo técnico
func (tr *TechnicianRepository) CreateNewTechnician(technician models.Technician) error {
	tr.dbFake = append(tr.dbFake, technician)
	return nil
}

// Obtener todos los técnicos
func (tr *TechnicianRepository) GetAllTechnicians() ([]models.UserDataView, error) {
	var users []models.UserDataView
	err := tr.db.Where("role_code = ?", "3").Find(&users).Error
	return users, err
}

// Obtener un técnico por ID
func (tr *TechnicianRepository) GetTechnicianById(id string) (models.GetTechnicianRequest, error) {
	for _, t := range tr.dbFake {
		if t.ID == id {
			technician := models.GetTechnicianRequest(t)
			return technician, nil
		}
	}
	return models.GetTechnicianRequest{}, fmt.Errorf("technician not found")
}

// Actualizar técnico por ID
func (tr *TechnicianRepository) UpdateTechnician(id string, updated models.Technician) error {
	for i, t := range tr.dbFake {
		if t.ID == id {
			updated.ID = id
			tr.dbFake[i].DocumentType = updated.DocumentType
			tr.dbFake[i].Document = updated.Document
			tr.dbFake[i].Name = updated.Name
			tr.dbFake[i].Surname = updated.Surname
			tr.dbFake[i].Email = updated.Email
			tr.dbFake[i].PhoneNumber = updated.PhoneNumber
			tr.dbFake[i].Username = updated.Username
			tr.dbFake[i].Address = updated.Address
			tr.dbFake[i].IsActive = updated.IsActive
			tr.dbFake[i].EntryDate = updated.EntryDate
			return nil
		}
	}
	return fmt.Errorf("technician not found")
}

// Eliminar técnico por ID
func (tr *TechnicianRepository) DeleteTechnician(id string) error {
	for i, t := range tr.dbFake {
		if t.ID == id {
			tr.dbFake = append(tr.dbFake[:i], tr.dbFake[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("technician not found")
}

// Cambiar contraseña de un técnico por Username
func (tr *TechnicianRepository) ChangePassword(username, newPassword string) error {
	for i, t := range tr.dbFake {
		if t.Username == username {
			tr.dbFake[i].Password = newPassword
			return nil
		}
	}
	return fmt.Errorf("technician with username %s not found", username)
}

// FindByUsername busca por Username o Email
func (tr *TechnicianRepository) FindByUsername(usernameOrEmail string) (*models.Technician, error) {
	for _, t := range tr.dbFake {
		if t.Username == usernameOrEmail || t.Email == usernameOrEmail {
			return &t, nil
		}
	}
	return nil, fmt.Errorf("technician not found")
}
