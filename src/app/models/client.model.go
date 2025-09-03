package models

import "github.com/google/uuid"

type Client struct {
	Person  Person
	Address string `json:"Address"`
}

type CreateClientRequest struct {
	DocumentType string `json:"DocumentType" validate:"required"`
	Document     string `json:"Document" validate:"required,numeric"`
	Name         string `json:"Name" validate:"required"`
	Surname      string `json:"Surname" validate:"required"`
	Email        string `json:"Email" validate:"required,email"`
	PhoneNumber  string `json:"PhoneNumber" validate:"required,numeric"`
	Address      string `json:"Address" validate:"required"`
}

type GetClientRequest struct {
	ID           uuid.UUID `json:"ID"`
	DocumentType string    `json:"DocumentType"`
	Document     string    `json:"Document"`
	Name         string    `json:"Name"`
	Surname      string    `json:"Surname"`
	Email        string    `json:"Email"`
	PhoneNumber  string    `json:"PhoneNumber"`
	Address      string    `json:"Address"`
}

type UpdateClientRequest struct {
	DocumentType string `json:"DocumentType" validate:"omitempty"`
	Document     string `json:"Document" validate:"omitempty"`
	Name         string `json:"Name" validate:"omitempty"`
	Surname      string `json:"Surname" validate:"omitempty"`
	Email        string `json:"Email" validate:"omitempty"`
	PhoneNumber  string `json:"PhoneNumber" validate:"omitempty"`
	Address      string `json:"Address" validate:"omitempty"`
}

func NewClient(documentType string, document string, name string, surname string, email string, phoneNumber string, address string) *Client {
	return &Client{
		Person: Person{
			ID:           uuid.New(),
			DocumentType: documentType,
			Document:     document,
			Name:         name,
			Surname:      surname,
			Email:        email,
			PhoneNumber:  phoneNumber,
		},
		Address: address,
	}
}
