package models

import (
	"time"

	"github.com/google/uuid"
)

type Technician struct {
	ID           uuid.UUID `json:"ID"`
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
