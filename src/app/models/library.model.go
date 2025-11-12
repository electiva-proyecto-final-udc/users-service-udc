package models

type LibraryModel struct {
	ID          string `json:"ID" gorm:"column:id"`
	Name        string `json:"Name" example:"Juan" gorm:"column:name"`
	Code        string `json:"Code" gorm:"column:code"`
	Description string `json:"Description" gorm:"column:description"`
}

func (LibraryModel) TableName() string{
	return "document_type"
}