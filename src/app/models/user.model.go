package models

import "encoding/json"

//  ENTITIES

type UserDataView struct {
	ID                      string  `json:"id" gorm:"column:id"`
	DocumentTypeName        string  `json:"documentTypeName" gorm:"column:document_type_name"`
	DocumentTypeDescription string  `json:"documentTypeDescription" gorm:"column:document_type_description"`
	DocumentCode            string  `json:"documentCode" gorm:"column:document_code"`
	DocumentNumber          string  `json:"documentNumber" gorm:"column:document_number"`
	Role                    string  `json:"role" gorm:"column:role"`
	RoleCode                string  `json:"roleCode" gorm:"column:role_code"`
	Name                    string  `json:"name" gorm:"column:name"`
	Surname                 string  `json:"surname" gorm:"column:surname"`
	Username                string  `json:"username" gorm:"column:username"`
	Email                   string  `json:"email" gorm:"column:email"`
	PhoneNumber             string  `json:"phoneNumber" gorm:"column:phone_number"`
	Address                 string  `json:"address" gorm:"column:address"`
	IsActive                bool    `json:"isActive" gorm:"column:isactive"`
	Permissions             *string `json:"permissions" gorm:"column:permissions"`
	UserId                  string  `json:"userId" gorm:"column:user_id"`
}

type UserProfileEntity struct {
	ID              string          `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username        string          `json:"username" gorm:"column:username"`
	Password        string          `json:"password" gorm:"column:password"`
	Permissions     json.RawMessage `json:"permissions" gorm:"type:jsonb;column:permissions"`
	PersonProfileID string          `json:"personProfileId" gorm:"column:person_profile_id;constraint:OnDelete:CASCADE;"`
}

func (UserDataView) TableName() string {
	return "user_data_view"
}

func (UserProfileEntity) TableName() string {
	return "user_profile"
}
