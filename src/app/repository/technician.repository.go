package repository

import (
	"fmt"
	"user-service-ucd/src/app/models"
)

// Se define la estructura del repositorio
type TechnicianRepository struct {
	db []models.Technician
}

// Constructor: se crea una nueva instancia
func NewTechnicianRepository() *TechnicianRepository {
	return &TechnicianRepository{
		db: []models.Technician{},
	}
}

// Crear un nuevo técnico
func (tr *TechnicianRepository) CreateNewTechnician(technician models.Technician) error {
	tr.db = append(tr.db, technician)
	return nil
}

// Obtener todos los técnicos 
func (tr *TechnicianRepository) GetAllTechnicians() ([]models.GetTechnicianRequest, error) {
	technicians := make([]models.GetTechnicianRequest, 0, len(tr.db))

	for _, t := range tr.db {
		technicians = append(technicians, models.GetTechnicianRequest(t))
	}

	return technicians, nil
}

// Obtener un técnico por ID
func (tr *TechnicianRepository) GetTechnicianById(id string) (models.GetTechnicianRequest, error) {
	for _, t := range tr.db {
		if t.ID == id {
			technician := models.GetTechnicianRequest(t)
			return technician, nil
		}
	}
	return models.GetTechnicianRequest{}, fmt.Errorf("technician not found")
}

// Actualizar técnico por ID
func (tr *TechnicianRepository) UpdateTechnician(id string, updated models.Technician) error {
	for i, t := range tr.db {
		if t.ID == id {
			updated.ID = id
			tr.db[i].DocumentType = updated.DocumentType
			tr.db[i].Document = updated.Document
			tr.db[i].Name = updated.Name
			tr.db[i].Surname = updated.Surname
			tr.db[i].Email = updated.Email
			tr.db[i].PhoneNumber = updated.PhoneNumber
			tr.db[i].Username = updated.Username
			tr.db[i].Address = updated.Address
			tr.db[i].IsActive = updated.IsActive
			tr.db[i].EntryDate = updated.EntryDate
			return nil
		}
	}
	return fmt.Errorf("technician not found")
}

// Eliminar técnico por ID
func (tr *TechnicianRepository) DeleteTechnician(id string) error {
	for i, t := range tr.db {
		if t.ID == id {
			tr.db = append(tr.db[:i], tr.db[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("technician not found")
}

// Cambiar contraseña de un técnico por Username
func (tr *TechnicianRepository) ChangePassword(username, newPassword string) error {
	for i, t := range tr.db {
		if t.Username == username {
			tr.db[i].Password = newPassword
			return nil
		}
	}
	return fmt.Errorf("technician with username %s not found", username)
}

// FindByUsername busca por Username o Email
func (tr *TechnicianRepository) FindByUsername(usernameOrEmail string) (*models.Technician, error) {
	for _, t := range tr.db {
		if t.Username == usernameOrEmail || t.Email == usernameOrEmail {
			return &t, nil
		}
	}
	return nil, fmt.Errorf("technician not found")
}