package dto

import (
	"encoding/json"
	"time"
)

type TechnicianDTO struct {
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

type CreateTechnicianDTO struct {
	DocumentTypeId string          `json:"DocumentTypeId" validate:"required" example:"61898ef1-914a-4435-a122-5f26fe253b1a"`
	RoleId         string          `json:"RoleId" validate:"required" example:"d290f1ee-6c54-4b01-90e6-d701748f0851"`
	Document       string          `json:"Document" validate:"required,numeric" example:"1234567890"`
	Name           string          `json:"Name" validate:"required" example:"Carlos"`
	Surname        string          `json:"Surname" validate:"required" example:"Ramírez"`
	Email          string          `json:"Email" validate:"required,email" example:"carlos.ramirez@example.com"`
	PhoneNumber    string          `json:"PhoneNumber" validate:"required,numeric" example:"3001234567"`
	Username       string          `json:"Username" validate:"required" example:"carlos.r"`
	Password       string          `json:"Password" validate:"required" example:"secret123"`
	Address        string          `json:"Address" validate:"required" example:"Calle 123 #45-67"`
	IsActive       bool            `json:"IsActive" validate:"required" example:"true"`
	EntryDate      time.Time       `json:"EntryDate" validate:"required" example:"2025-01-01T15:04:05Z"`
	Permissions    json.RawMessage `json:"Permissions" validate:"required" example:"{}"`
}

type ChangePasswordDTO struct {
	UserId      string `json:"UserId" validate:"required" example:"carlos.r"`
	NewPassword string `json:"NewPassword" validate:"required" example:"newSecret123"`
}

type GetTechnicianDTO struct {
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

type UpdateTechnicianDTO struct {
	DocumentTypeId string `json:"DocumentTypeId" validate:"required" example:"CC"`
	Document       string `json:"Document" validate:"required" example:"9876543210"`
	Name           string `json:"Name" validate:"required" example:"Juan"`
	Email          string `json:"Email" validate:"required" example:"juan.perez@example.com"`
	PhoneNumber    string `json:"PhoneNumber" validate:"required" example:"3017654321"`
	Address        string `json:"Address" validate:"required" example:"Carrera 45 #12-34"`
}
