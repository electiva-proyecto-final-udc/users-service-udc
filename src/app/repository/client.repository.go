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
			ID:           c.ID,
			DocumentType: c.DocumentType,
			Document:     c.Document,
			Name:         c.Name,
			Surname:      c.Surname,
			Email:        c.Email,
			PhoneNumber:  c.PhoneNumber,
			Address:      c.Address,
		})
	}

	return clients, nil
}

func (cr *ClientRepository) GetClientById(id uuid.UUID) (models.GetClientRequest, error) {
	for _, c := range cr.db {
		if c.ID == id {
			client := models.GetClientRequest{
				ID:           c.ID,
				DocumentType: c.DocumentType,
				Document:     c.Document,
				Name:         c.Name,
				Surname:      c.Surname,
				Email:        c.Email,
				PhoneNumber:  c.PhoneNumber,
				Address:      c.Address}
			return client, nil
		}
	}

	return models.GetClientRequest{}, fmt.Errorf("client not found")
}

func (cr *ClientRepository) UpdateClient(id uuid.UUID, updated models.Client) error {
	for i, c := range cr.db {
		if c.ID == id {
			updated.ID = id
			cr.db[i].DocumentType = updated.DocumentType
			cr.db[i].Document     = updated.Document
			cr.db[i].Name         = updated.Name
			cr.db[i].Surname      = updated.Surname
			cr.db[i].Email        = updated.Email
			cr.db[i].PhoneNumber  = updated.PhoneNumber
			cr.db[i].Address = updated.Address
			return nil
		}
	}
	return fmt.Errorf("client not found")
}

func (cr *ClientRepository) DeleteClient(id uuid.UUID) error {
	for i, c := range cr.db {
		if c.ID == id {
			cr.db = append(cr.db[:i], cr.db[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("client not found")
}
