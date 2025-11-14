package services

import (
	"fmt"
	"user-service-ucd/src/app/dto"
	"user-service-ucd/src/app/models"
	"user-service-ucd/src/app/repository"

	"golang.org/x/crypto/bcrypt"
)

type TechnicianService struct {
	tr *repository.TechnicianRepository
	ns *NotificationService
}

// Constructor
func NewTechnicianService(tr *repository.TechnicianRepository, ns *NotificationService) *TechnicianService {
	return &TechnicianService{
		tr: tr,
		ns: ns,
	}
}

// Crear nuevo técnico
func (ts *TechnicianService) NewTechnician(technician dto.CreateTechnicianDTO, token string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(technician.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	
	// Envía el correo al usuario al registrarse
	msg, errSendingEmail := ts.ns.SendWelcomeEmail(technician.Email, dto.LoginRequest{
		Username: technician.Username,
		Password: technician.Password,
	}, token)
	fmt.Println(msg)
	if errSendingEmail != nil {
		fmt.Println(errSendingEmail)
	}

	technician.Password = string(hashedPassword)
	response := ts.tr.CreateNewTechnician(technician)
	if response != nil {
		return response
	}

	
	return response
}

// Obtener todos los técnicos
func (ts *TechnicianService) GetAllTechnicians() ([]models.UserDataView, error) {
	return ts.tr.GetAllTechnicians()
}

// Obtener técnico por ID
func (ts *TechnicianService) GetTechnicianById(id string) (models.UserDataView, error) {
	return ts.tr.GetTechnicianById(id)
}

// Actualizar técnico
func (ts *TechnicianService) UpdateTechnician(id string, technician dto.UpdateTechnicianDTO) error {
	return ts.tr.UpdateTechnician(id, technician)
}

// Eliminar técnico
func (ts *TechnicianService) DeleteTechnician(id string) error {
	return ts.tr.DeleteTechnician(id)
}

// Cambiar contraseña
func (ts *TechnicianService) ChangePassword(changePassWordRequest dto.ChangePasswordDTO) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(changePassWordRequest.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	changePassWordRequest.NewPassword = string(hashedPassword)
	return ts.tr.ChangePassword(changePassWordRequest)
}
