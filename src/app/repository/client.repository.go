package repository

import (
	"encoding/json"
	"fmt"
	"user-service-ucd/src/app/dto"
	"user-service-ucd/src/app/models"
	"user-service-ucd/src/database"

	"gorm.io/gorm"
)

// Se define la estructura del repositorio
type ClientRepository struct {
	db     *gorm.DB
}

// Se crea una nueva instancia (Constructor)
func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{
		db:     database.DB,
	}
}

func (cr *ClientRepository) CreateNewClient(client dto.CreateClientRequest) (error) {
	var personProfileData models.PersonProfile
	var role models.RoleModel
	errRole := cr.db.Where("description = ?", "CLIENT").Find(&role).Error

	if errRole != nil {
		return errRole
	}
	client.RoleId = role.ID
	clientData, _ := json.Marshal(client)

	if err := json.Unmarshal(clientData, &personProfileData); err != nil {
		return err
	}

	if err := cr.db.Create(&personProfileData).Error; err != nil {
		return err
	}

	return nil
}

// Obtener todos los clientes transformados a GetClientRequest
func (cr *ClientRepository) GetAllClients() ([]models.ClientDataView, error) {
	var clients []models.ClientDataView
	err := cr.db.Find(&clients).Error
	return clients, err
}

func (cr *ClientRepository) GetClientById(id string) (models.ClientDataView, error) {
	var client models.ClientDataView
	err := cr.db.Where("id = ?", id).Find(&client).Error

	if err != nil {
		return models.ClientDataView{}, fmt.Errorf("ERROR FETCHING CLIENT")
	}

	if client == (models.ClientDataView{}) {
		return models.ClientDataView{}, fmt.Errorf("CLIENT NOT FOUND")
	}

	return client, nil
}

func (cr *ClientRepository) GetClientByDocument(id string) (models.ClientDataView, error) {
	var client models.ClientDataView
	err := cr.db.Where("document_number = ?", id).Find(&client).Error

	if err != nil {
		return models.ClientDataView{}, fmt.Errorf("ERROR FETCHING CLIENT")
	}

	if client == (models.ClientDataView{}) {
		return models.ClientDataView{}, fmt.Errorf("CLIENT NOT FOUND")
	}

	return client, nil
}

func (cr *ClientRepository) UpdateClient(id string, updated dto.UpdateClientRequest) error {
	var clientUpdated models.UpdateClientEntity
	clientData, _:= json.Marshal(updated)
	if err := json.Unmarshal(clientData, &clientUpdated); err != nil{
		return err
	}

	result := cr.db.Model(&models.UpdateClientEntity{}).Where("id = ?", id).Updates(clientUpdated)
	
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("CLIENT NOT FOUND")
	}

	return nil
}

func (cr *ClientRepository) DeleteClient(id string) error {
	result := cr.db.Model(models.PersonProfile{}).Where("id = ?", id).Delete(nil)

	if result.Error != nil {
		return fmt.Errorf("ERROR DELETING CLIENT")
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("CLIENT NOT FOUND")
	}

	return nil
}
