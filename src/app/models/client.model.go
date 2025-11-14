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

type ClientDataView struct {
	ID                      string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" gorm:"column:id"`
	Document                string `json:"document" example:"1002003001" gorm:"column:document_number"`
	DocumentTypeName        string `json:"document_type_name" gorm:"column:document_type_name"`
	DocumentTypeDescription string `json:"document_type_description" gorm:"column:document_type_description"`
	Role                    string `json:"role" gorm:"column:role"`
	Name                    string `json:"name" example:"Juan" gorm:"column:name"`
	Surname                 string `json:"surname" example:"Pérez" gorm:"column:surname"`
	Email                   string `json:"email" example:"juan.perez@example.com" gorm:"column:email"`
	PhoneNumber             string `json:"phoneNumber" example:"3001234567" gorm:"column:phone_number"`
	Address                 string `json:"address" example:"Calle 123 #45-67" gorm:"column:address"`
}

func (ClientDataView) TableName() string {
	return "client_data_view"
}

type UpdateClientEntity struct {
	DocumentTypeId string `json:"DocumentTypeId" example:"CC" gorm:"column:document_type_id"`
	Document       string `json:"Document" example:"1002003001" gorm:"column:document"`
	Name           string `json:"Name" example:"Juan" gorm:"column:name"`
	Surname        string `json:"Surname" example:"Pérez" gorm:"column:surname"`
	Email          string `json:"Email" example:"juan.perez@example.com" gorm:"column:email"`
	PhoneNumber    string `json:"PhoneNumber" example:"3001234567" gorm:"column:phone_number"`
	Address        string `json:"Address" example:"Calle 123 #45-67" gorm:"column:address"`
}

func (UpdateClientEntity) TableName() string {
	return "person_profile"
}
