package models

type DocumentTypeModel struct {
	ID          string `json:"id" gorm:"column:id"`
	Name        string `json:"name" example:"Juan" gorm:"column:name"`
	Code        string `json:"code" gorm:"column:code"`
	Description string `json:"description" gorm:"column:description"`
}

func (DocumentTypeModel) TableName() string{
	return "document_type"
}

type RoleModel struct {
	ID          string `json:"id" gorm:"column:id"`
	Code        string `json:"dode" gorm:"column:code"`
	Description string `json:"description" gorm:"column:description"`
}

func (RoleModel) TableName() string {
	return "role"
}