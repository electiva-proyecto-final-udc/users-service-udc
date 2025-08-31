package repository

import (
	"user-service-ucd/src/app/models"
)

// Se define la estructura del repositorio
type ClientRepository struct {
	db []models.Client
}

// Se crea una nueva instancia (Constructor)
func NewClientRepository() *ClientRepository {
	return &ClientRepository{
		db: []models.Client{},
	}
}

func (cr *ClientRepository) CreateNewClient(client models.Client) {
	cr.db = append(cr.db, client)
}

// Obtener todos los clientes transformados a GetClientRequest
func (cr *ClientRepository) GetAllClients() []models.GetClientRequest {
	clients := make([]models.GetClientRequest, 0, len(cr.db))

	for _, c := range cr.db {
		clients = append(clients, models.GetClientRequest{
			ID:           c.Person.ID,
			DocumentType: c.Person.DocumentType,
			Document:     c.Person.Document,
			Name:         c.Person.Name,
			Surname:      c.Person.Surname,
			Email:        c.Person.Email,
			PhoneNumber:  c.Person.PhoneNumber,
			Address:      c.Address,
		})
	}

	return clients
}
