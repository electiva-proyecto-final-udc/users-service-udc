package services

import (
	"user-service-ucd/src/app/models"
	"user-service-ucd/src/app/repository"
)

type LibraryService struct {
	lr *repository.LibraryRepository
}

func NewLibraryService(lr *repository.LibraryRepository) *LibraryService{
	return &LibraryService{
		lr: lr,
	}
}

func (ls *LibraryService) GetDocumentTypes() ([]models.LibraryModel, error){
	return ls.lr.GetDocumentTypes()
}