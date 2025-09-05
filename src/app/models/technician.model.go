package models

import (
	"time"
	"user-service-ucd/utils"

	"github.com/google/uuid"
)

type Technician struct {
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

type CreateTechnicianRequest struct {
	DocumentType string    `json:"DocumentType" validate:"required"`
	Document     string    `json:"Document" validate:"required,numeric"`
	Name         string    `json:"Name" validate:"required"`
	Surname      string    `json:"Surname" validate:"required"`
	Email        string    `json:"Email" validate:"required,email"`
	PhoneNumber  string    `json:"PhoneNumber" validate:"required,numeric"`
	Username     string    `json:"Username" validate:"required"`
	Password     string    `json:"Password" validate:"required"`
	Address      string    `json:"Address" validate:"required"`
	IsActive     bool      `json:"IsActive" validate:"required"`
	EntryDate    time.Time `json:"EntryDate" validate:"required"`
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

type UpdateTechnicianRequest struct {
	DocumentType string    `json:"DocumentType"`
	Document     string    `json:"Document"`
	Name         string    `json:"Name"`
	Surname      string    `json:"Surname"`
	Email        string    `json:"Email"`
	PhoneNumber  string    `json:"PhoneNumber"`
	Username     string    `json:"Username"`
	Address      string    `json:"Address"`
	IsActive     bool      `json:"IsActive"`
	EntryDate    time.Time `json:"EntryDate"`
}

type ChangePasswordRequest struct {
	Username    string `json:"Username"`
	NewPassword string `json:"NewPassword"`
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
		Email:        technician.Email,
		PhoneNumber:  technician.PhoneNumber,
		Address:      technician.Address,
	}, nil
}
