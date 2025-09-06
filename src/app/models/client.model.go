package models

import "github.com/google/uuid"

type Client struct {
	ID           uuid.UUID `json:"ID" example:"550e8400-e29b-41d4-a716-446655440000"`
	DocumentType string    `json:"DocumentType" example:"CC"`
	Document     string    `json:"Document" example:"1002003001"`
	Name         string    `json:"Name" example:"Juan"`
	Surname      string    `json:"Surname" example:"Pérez"`
	Email        string    `json:"Email" example:"juan.perez@example.com"`
	PhoneNumber  string    `json:"PhoneNumber" example:"3001234567"`
	Address      string    `json:"Address" example:"Calle 123 #45-67"`
}

type CreateClientRequest struct {
	DocumentType string `json:"DocumentType" validate:"required" example:"CC"`
	Document     string `json:"Document" validate:"required,numeric" example:"1002003001"`
	Name         string `json:"Name" validate:"required" example:"Juan"`
	Surname      string `json:"Surname" validate:"required" example:"Pérez"`
	Email        string `json:"Email" validate:"required,email" example:"juan.perez@example.com"`
	PhoneNumber  string `json:"PhoneNumber" validate:"required,numeric" example:"3001234567"`
	Address      string `json:"Address" validate:"required" example:"Calle 123 #45-67"`
}

type GetClientRequest struct {
	ID           uuid.UUID `json:"ID" example:"550e8400-e29b-41d4-a716-446655440000"`
	DocumentType string    `json:"DocumentType" example:"CC"`
	Document     string    `json:"Document" example:"1002003001"`
	Name         string    `json:"Name" example:"Juan"`
	Surname      string    `json:"Surname" example:"Pérez"`
	Email        string    `json:"Email" example:"juan.perez@example.com"`
	PhoneNumber  string    `json:"PhoneNumber" example:"3001234567"`
	Address      string    `json:"Address" example:"Calle 123 #45-67"`
}

type UpdateClientRequest struct {
	DocumentType string `json:"DocumentType" validate:"omitempty" example:"CC"`
	Document     string `json:"Document" validate:"omitempty" example:"1002003001"`
	Name         string `json:"Name" validate:"omitempty" example:"Juan"`
	Surname      string `json:"Surname" validate:"omitempty" example:"Pérez"`
	Email        string `json:"Email" validate:"omitempty" example:"juan.perez@example.com"`
	PhoneNumber  string `json:"PhoneNumber" validate:"omitempty" example:"3001234567"`
	Address      string `json:"Address" validate:"omitempty" example:"Calle 123 #45-67"`
}

func NewClient(documentType string, document string, name string, surname string, email string, phoneNumber string, address string) *Client {
	return &Client{
		ID:           uuid.New(),
		DocumentType: documentType,
		Document:     document,
		Name:         name,
		Surname:      surname,
		Email:        email,
		PhoneNumber:  phoneNumber,
		Address: address,
	}
}
