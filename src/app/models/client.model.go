package models

import "github.com/google/uuid"

type Client struct {
	Person  Person
	Address string `json:"Address"`
}

type CreateClientRequest struct {
	DocumentType string `json:"DocumentType"`
	Document     string `json:"Document"`
	Name         string `json:"Name"`
	Surname      string `json:"Surname"`
	Email        string `json:"Email"`
	PhoneNumber  string `json:"PhoneNumber"`
	Address      string `json:"Address"`
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
