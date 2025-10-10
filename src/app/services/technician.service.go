package services

import (
	"user-service-ucd/src/app/models"
	"user-service-ucd/src/app/repository"

	"golang.org/x/crypto/bcrypt"
)

type TechnicianService struct {
	tr *repository.TechnicianRepository
}

// Constructor
func NewTechnicianService(tr *repository.TechnicianRepository) *TechnicianService {
	return &TechnicianService{
		tr: tr,
	}
}

// Crear nuevo técnico
func (ts *TechnicianService) NewTechnician(technician models.Technician) error {
	return ts.tr.CreateNewTechnician(technician)
}

// Obtener todos los técnicos
func (ts *TechnicianService) GetAllTechnicians() ([]models.UserDataView, error) {
	return ts.tr.GetAllTechnicians()
}

// Obtener técnico por ID
func (ts *TechnicianService) GetTechnicianById(id string) (models.GetTechnicianRequest, error) {
	return ts.tr.GetTechnicianById(id)
}

// Actualizar técnico
func (ts *TechnicianService) UpdateTechnician(id string, technician models.Technician) error {
	return ts.tr.UpdateTechnician(id, technician)
}

// Eliminar técnico
func (ts *TechnicianService) DeleteTechnician(id string) error {
	return ts.tr.DeleteTechnician(id)
}

// Cambiar contraseña
func (ts *TechnicianService) ChangePassword(username, newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return ts.tr.ChangePassword(username, string(hashedPassword))
}
