package models

type UpdateTechnician struct {
	DocumentTypeId string `json:"DocumentTypeId" example:"CC" gorm:"column:document_type_id"`
	Document       string `json:"Document" example:"9876543210" gorm:"column:document"`
	Name           string `json:"Name" example:"Juan" gorm:"column:name"`
	Email          string `json:"Email" example:"juan.perez@example.com" gorm:"column:email"`
	PhoneNumber    string `json:"PhoneNumber" example:"3017654321" gorm:"column:phone_number"`
	Address        string `json:"Address" example:"Carrera 45 #12-34" gorm:"column:address"`
}

// TABLAS
func (UpdateTechnician) TableName() string {
	return "person_profile"
}
