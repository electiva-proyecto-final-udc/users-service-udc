package models

type UpdateTechnician struct {
	DocumentTypeId string `json:"documentTypeId" example:"CC" gorm:"column:document_type_id"`
	Document       string `json:"document" example:"9876543210" gorm:"column:document"`
	Name           string `json:"name" example:"Juan" gorm:"column:name"`
	Email          string `json:"email" example:"juan.perez@example.com" gorm:"column:email"`
	PhoneNumber    string `json:"phoneNumber" example:"3017654321" gorm:"column:phone_number"`
	Address        string `json:"address" example:"Carrera 45 #12-34" gorm:"column:address"`
}

// TABLAS
func (UpdateTechnician) TableName() string {
	return "person_profile"
}
