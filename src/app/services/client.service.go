package services

import (
	"user-service-ucd/src/app/models"
	"user-service-ucd/src/app/repository"
)

type ClientService struct {
	cr *repository.ClientRepository
}

func NewClientService(cr *repository.ClientRepository) *ClientService{
	return &ClientService{
		cr: cr,
	}
}

func (cs *ClientService) NewClient(client models.Client){
	cs.cr.CreateNewClient(client)
}

func (cs *ClientService) GetAllClients() []models.GetClientRequest{
	return cs.cr.GetAllClients()
}