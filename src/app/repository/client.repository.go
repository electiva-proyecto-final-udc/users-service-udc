package repository

import (
	"fmt"
	"user-service-ucd/src/app/models"

	"github.com/google/uuid"
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

func (cr *ClientRepository) CreateNewClient(client models.Client) error {
	cr.db = append(cr.db, client)
	return nil
}

// Obtener todos los clientes transformados a GetClientRequest
func (cr *ClientRepository) GetAllClients() ([]models.GetClientRequest, error) {
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

	return clients, nil
}

func (cr *ClientRepository) GetClientById(id uuid.UUID) (models.GetClientRequest, error) {
	for _, c := range cr.db {
		if c.Person.ID == id {
			client := models.GetClientRequest{
				ID:           c.Person.ID,
				DocumentType: c.Person.DocumentType,
				Document:     c.Person.Document,
				Name:         c.Person.Name,
				Surname:      c.Person.Surname,
				Email:        c.Person.Email,
				PhoneNumber:  c.Person.PhoneNumber,
				Address:      c.Address}
			return client, nil
		}
	}

	return models.GetClientRequest{}, fmt.Errorf("client not found")
}

func (cr *ClientRepository) UpdateClient(id uuid.UUID, updated models.Client) error {
	for i, c := range cr.db {
		if c.Person.ID == id {
			updated.Person.ID = id
			cr.db[i].Person.DocumentType = updated.Person.DocumentType
			cr.db[i].Person.Document     = updated.Person.Document
			cr.db[i].Person.Name         = updated.Person.Name
			cr.db[i].Person.Surname      = updated.Person.Surname
			cr.db[i].Person.Email        = updated.Person.Email
			cr.db[i].Person.PhoneNumber  = updated.Person.PhoneNumber
			cr.db[i].Address = updated.Address
			return nil
		}
	}
	return fmt.Errorf("client not found")
}

func (cr *ClientRepository) DeleteClient(id uuid.UUID) error {
	for i, c := range cr.db {
		if c.Person.ID == id {
			cr.db = append(cr.db[:i], cr.db[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("client not found")
}
