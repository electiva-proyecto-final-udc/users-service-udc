package models

import (
	"time"
	"user-service-ucd/utils"

	"github.com/google/uuid"
)

type Technician struct {
	ID           string    `json:"ID" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
	DocumentType string    `json:"DocumentType" example:"CC"`
	Document     string    `json:"Document" example:"1234567890"`
	Name         string    `json:"Name" example:"Carlos"`
	Surname      string    `json:"Surname" example:"Ramírez"`
	Email        string    `json:"Email" example:"carlos.ramirez@example.com"`
	PhoneNumber  string    `json:"PhoneNumber" example:"3001234567"`
	Username     string    `json:"Username" example:"carlos.r"`
	Password     string    `json:"Password,omitempty" example:"$2a$10$hashPassword"`
	Address      string    `json:"Address" example:"Calle 123 #45-67"`
	IsActive     bool      `json:"IsActive" example:"true"`
	EntryDate    time.Time `json:"EntryDate" example:"2025-01-01T15:04:05Z"`
}

type CreateTechnicianRequest struct {
	DocumentType string    `json:"DocumentType" validate:"required" example:"CC"`
	Document     string    `json:"Document" validate:"required,numeric" example:"1234567890"`
	Name         string    `json:"Name" validate:"required" example:"Carlos"`
	Surname      string    `json:"Surname" validate:"required" example:"Ramírez"`
	Email        string    `json:"Email" validate:"required,email" example:"carlos.ramirez@example.com"`
	PhoneNumber  string    `json:"PhoneNumber" validate:"required,numeric" example:"3001234567"`
	Username     string    `json:"Username" validate:"required" example:"carlos.r"`
	Password     string    `json:"Password" validate:"required" example:"secret123"`
	Address      string    `json:"Address" validate:"required" example:"Calle 123 #45-67"`
	IsActive     bool      `json:"IsActive" validate:"required" example:"true"`
	EntryDate    time.Time `json:"EntryDate" validate:"required" example:"2025-01-01T15:04:05Z"`
}

type UpdateTechnicianRequest struct {
	DocumentType string    `json:"DocumentType" example:"CC"`
	Document     string    `json:"Document" example:"9876543210"`
	Name         string    `json:"Name" example:"Juan"`
	Surname      string    `json:"Surname" example:"Pérez"`
	Email        string    `json:"Email" example:"juan.perez@example.com"`
	PhoneNumber  string    `json:"PhoneNumber" example:"3017654321"`
	Username     string    `json:"Username" example:"juan.p"`
	Address      string    `json:"Address" example:"Carrera 45 #12-34"`
	IsActive     bool      `json:"IsActive" example:"false"`
	EntryDate    time.Time `json:"EntryDate" example:"2025-02-01T10:00:00Z"`
}

type ChangePasswordRequest struct {
	Username    string `json:"Username" example:"carlos.r"`
	NewPassword string `json:"NewPassword" example:"newSecret123"`
}

type GetTechnicianRequest struct {
	ID           string    `json:"ID"`
	DocumentType string    `json:"DocumentType"`
	Document     string    `json:"Document"`
	Name         string    `json:"Name"`
	Surname      string    `json:"Surname"`
	Email        string    `json:"Email"`
	PhoneNumber  string    `json:"PhoneNumber"`
	Username     string    `json:"Username"`
	Password     string    `json:"Password"`
	Address      string    `json:"Address"`
	IsActive     bool      `json:"IsActive"`
	EntryDate    time.Time `json:"EntryDate"`
}

func NewTechnician(technician CreateTechnicianRequest) (*Technician, error) {
	hashedPassword, err := utils.HashPassword(technician.Password)
	if err != nil {
		return nil, err
	}
	return &Technician{
		ID:           uuid.New().String(),
		DocumentType: technician.DocumentType,
		Document:     technician.Document,
		Name:         technician.Name,
		Surname:      technician.Surname,
		Email:        technician.Email,
		PhoneNumber:  technician.PhoneNumber,
		Address:      technician.Address,
		Username:     technician.Username,
		Password:     hashedPassword,
		EntryDate:    technician.EntryDate,
	}, nil
}

func NewTechnicianUpdated(technician UpdateTechnicianRequest, id string) (*Technician, error) {
	return &Technician{
		ID:           id,
		DocumentType: technician.DocumentType,
		Document:     technician.Document,
		Name:         technician.Name,
		Surname:      technician.Surname,
		Username:     technician.Username,
		Email:        technician.Email,
		PhoneNumber:  technician.PhoneNumber,
		Address:      technician.Address,
	}, nil
}
