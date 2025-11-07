package services

import (
	"user-service-ucd/src/app/dto"
	"user-service-ucd/src/app/models"
	"user-service-ucd/src/app/repository"
)

type ClientService struct {
	cr *repository.ClientRepository
}

func NewClientService(cr *repository.ClientRepository) *ClientService {
	return &ClientService{
		cr: cr,
	}
}

func (cs *ClientService) NewClient(client dto.CreateClientRequest) error {
	return cs.cr.CreateNewClient(client)
}

func (cs *ClientService) GetAllClients() ([]models.ClientDataView, error) {
	return cs.cr.GetAllClients()
}

func (cs *ClientService) GetClientById(id string) (models.ClientDataView, error) {
	return cs.cr.GetClientById(id)
}

func (cs *ClientService) UpdateClient(id string, client dto.UpdateClientRequest) error {
	return cs.cr.UpdateClient(id, client)
}

func (cs *ClientService) DeleteClient(id string) error {
	return cs.cr.DeleteClient(id)
}
