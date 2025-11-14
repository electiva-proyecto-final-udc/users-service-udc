package dto

type CreateClientRequest struct {
	ID             string `json:"id" validate:"omitempty"`
	DocumentTypeId string `json:"documentTypeId" validate:"required" example:"CC"`
	RoleId         string `json:"roleId" validate:"omitempty" example:"CC"`
	Document       string `json:"document" validate:"required,numeric" example:"1002003001"`
	Name           string `json:"name" validate:"required" example:"Juan"`
	Surname        string `json:"surname" validate:"required" example:"Pérez"`
	Email          string `json:"email" validate:"required,email" example:"juan.perez@example.com"`
	PhoneNumber    string `json:"phoneNumber" validate:"required,numeric" example:"3001234567"`
	Address        string `json:"address" validate:"required" example:"Calle 123 #45-67"`
}

type UpdateClientRequest struct {
	DocumentTypeId string `json:"DocumentTypeId" validate:"omitempty" example:"CC"`
	Document       string `json:"Document" validate:"omitempty" example:"1002003001"`
	Name           string `json:"Name" validate:"omitempty" example:"Juan"`
	Surname        string `json:"Surname" validate:"omitempty" example:"Pérez"`
	Email          string `json:"Email" validate:"omitempty" example:"juan.perez@example.com"`
	PhoneNumber    string `json:"PhoneNumber" validate:"omitempty" example:"3001234567"`
	Address        string `json:"Address" validate:"omitempty" example:"Calle 123 #45-67"`
}
