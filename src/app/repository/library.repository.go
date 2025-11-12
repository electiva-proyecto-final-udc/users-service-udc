package repository

import (
	"user-service-ucd/src/app/models"
	"user-service-ucd/src/database"

	"gorm.io/gorm"
)

type LibraryRepository struct {
	db *gorm.DB
}

func NewLibraryRepository(db *gorm.DB) *LibraryRepository {
	return &LibraryRepository{
		db: database.DB,
	}
}

func (tr *LibraryRepository) GetDocumentTypes() ([]models.LibraryModel, error){
	var documentTypes []models.LibraryModel
	err := tr.db.Find(&documentTypes).Error
	return documentTypes, err
}