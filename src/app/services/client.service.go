package services

import (
	"user-service-ucd/src/app/models"
	"user-service-ucd/src/app/repository"

	"github.com/google/uuid"
)

type ClientService struct {
	cr *repository.ClientRepository
}

func NewClientService(cr *repository.ClientRepository) *ClientService {
	return &ClientService{
		cr: cr,
	}
}

func (cs *ClientService) NewClient(client models.Client) error {
	return cs.cr.CreateNewClient(client)
}

func (cs *ClientService) GetAllClients() ([]models.GetClientRequest, error) {
	return cs.cr.GetAllClients()
}

func (cs *ClientService) GetClientById(id uuid.UUID) (models.GetClientRequest, error) {
	return cs.cr.GetClientById(id)
}

func (cs *ClientService) UpdateClient(id uuid.UUID, client models.Client) (error) {
	return cs.cr.UpdateClient(id, client)
} 

func (cs *ClientService) DeleteClient(id uuid.UUID) (error) {
	return cs.cr.DeleteClient(id)
}