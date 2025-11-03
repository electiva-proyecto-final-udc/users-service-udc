package models

import "encoding/json"

// REQUESTS MODELS (DTO'S)

type LoginRequest struct {
	Username string `json:"Username" example:"carlos.r"`
	Password string `json:"Password" example:"secret123"`
}

type UserInfo struct {
	ID       string `json:"ID"`
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Role     string `json:"Role"`
}

//  ENTITIES

type UserDataView struct {
	ID                      string  `json:"id" gorm:"column:id"`
	DocumentTypeName        string  `json:"document_type_name" gorm:"column:document_type_name"`
	DocumentTypeDescription string  `json:"document_type_description" gorm:"column:document_type_description"`
	DocumentCode            string  `json:"document_code" gorm:"column:document_code"`
	DocumentNumber          string  `json:"document_number" gorm:"column:document_number"`
	Role                    string  `json:"role" gorm:"column:role"`
	RoleCode                string  `json:"role_code" gorm:"column:role_code"`
	Name                    string  `json:"name" gorm:"column:name"`
	Surname                 string  `json:"surname" gorm:"column:surname"`
	Email                   string  `json:"email" gorm:"column:email"`
	PhoneNumber             string  `json:"phone_number" gorm:"column:phone_number"`
	Address                 string  `json:"address" gorm:"column:address"`
	IsActive                bool    `json:"is_active" gorm:"column:isactive"`
	Permissions             *string `json:"permissions" gorm:"column:permissions"`
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
