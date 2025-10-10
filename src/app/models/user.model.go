package models

type UserDataView struct {
	ID                       string    `json:"id" gorm:"column:id"`
	DocumentTypeName         string  `json:"document_type_name" gorm:"column:document_type_name"`
	DocumentTypeDescription  string  `json:"document_type_description" gorm:"column:document_type_description"`
	DocumentCode             string  `json:"document_code" gorm:"column:document_code"`
	DocumentNumber           string  `json:"document_number" gorm:"column:document_number"`
	Role                     string  `json:"role" gorm:"column:role"`
	RoleCode                 string  `json:"role_code" gorm:"column:role_code"`
	Name                     string  `json:"name" gorm:"column:name"`
	Surname                  string  `json:"surname" gorm:"column:surname"`
	Email                    string  `json:"email" gorm:"column:email"`
	PhoneNumber              string  `json:"phone_number" gorm:"column:phone_number"`
	Address                  string  `json:"address" gorm:"column:address"`
	IsActive                 bool    `json:"is_active" gorm:"column:isactive"`
	Permissions              *string `json:"permissions" gorm:"column:permissions"`
}

func (UserDataView) TableName() string {
	return "user_data_view"
}