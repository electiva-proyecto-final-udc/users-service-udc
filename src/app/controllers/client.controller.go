package controllers

import (
	"encoding/json"
	"net/http"
	"user-service-ucd/src/app/models"
	"user-service-ucd/src/app/services"
	"user-service-ucd/src/common"
)

type ClientController struct {
	cs *services.ClientService
}

func NewClientController(cs *services.ClientService) *ClientController {
	return &ClientController{
		cs: cs,
	}
}

func (cc *ClientController) AddNewClient(w http.ResponseWriter, r *http.Request) {
	var request models.CreateClientRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.Header().Set("Content-Type", "application/json") // 400
		w.WriteHeader(http.StatusBadRequest)               // 400
		json.NewEncoder(w).Encode(common.ApiResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Error: &common.ErrorResponse{
				Message: "Invalid request body",
				Details: err.Error(),
			},
		})
		return
	}

	client := models.NewClient(
		request.DocumentType,
		request.Document,
		request.Name,
		request.Surname,
		request.Email,
		request.PhoneNumber,
		request.Address,
	)

	cc.cs.NewClient(*client)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(common.ApiResponse{
		Success:    true,
		StatusCode: http.StatusCreated,
		Message:    "CLIENT_CREATED_SUCCESFULLY",
		Data:       request,
	})
}

func (cc *ClientController) GetAllClients(w http.ResponseWriter, r *http.Request) {
	clients := cc.cs.GetAllClients()
	if len(clients) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(common.ApiResponse{
			Success:    true,
			StatusCode: http.StatusNoContent,
			Message: "NO_CONTENT",
			Data:       clients,
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(common.ApiResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "CLIENTS_FETCHED_SUCCESSFULLY",
		Data:       clients,
	})
}
