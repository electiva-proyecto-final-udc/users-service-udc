package models

type PersonProfile struct {
	ID             string `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	DocumentTypeId string `json:"documentTypeId" gorm:"column:document_type_id"`
	RoleId         string `json:"roleId" gorm:"column:role_id"`
	Name           string `json:"name" gorm:"column:name"`
	Surname        string `json:"surname" gorm:"column:surname"`
	Email          string `json:"email" gorm:"column:email"`
	PhoneNumber    string `json:"phoneNumber" gorm:"column:phone_number"`
	Address        string `json:"address" gorm:"column:address"`
	Document       string `json:"document" gorm:"column:document"`
}

func (PersonProfile) TableName() string {
	return "person_profile"
}
