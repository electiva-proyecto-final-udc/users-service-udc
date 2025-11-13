package models

type DocumentTypeModel struct {
	ID          string `json:"ID" gorm:"column:id"`
	Name        string `json:"Name" example:"Juan" gorm:"column:name"`
	Code        string `json:"Code" gorm:"column:code"`
	Description string `json:"Description" gorm:"column:description"`
}

func (DocumentTypeModel) TableName() string{
	return "document_type"
}

type RoleModel struct {
	ID          string `json:"ID" gorm:"column:id"`
	Code        string `json:"Code" gorm:"column:code"`
	Description string `json:"Description" gorm:"column:description"`
}

func (RoleModel) TableName() string {
	return "role"
}