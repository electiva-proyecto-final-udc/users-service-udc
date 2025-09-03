package controllers

import (
	"encoding/json"
	"net/http"
	"user-service-ucd/src/app/models"
	"user-service-ucd/src/app/services"
	"user-service-ucd/src/common"
	"user-service-ucd/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
		common.JSONResponse(w, http.StatusBadRequest, common.ApiResponse{
			Error: &common.ErrorResponse{
				Message: "Invalid request body",
				Details: err.Error(),
			},
		})
		return
	}

	if errors := utils.ValidateEntity(request); errors != nil {
		common.JSONResponse(w, http.StatusUnprocessableEntity, common.ApiResponse{
			Error: &common.ErrorResponse{
				Message: "Invalid data",
				Details: errors,
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
	common.JSONResponse(w, http.StatusCreated, common.ApiResponse{
		Message: "CLIENT_CREATED_SUCCESSFULLY",
		Data:    request,
	})
}

func (cc *ClientController) GetAllClients(w http.ResponseWriter, r *http.Request) {
	clients, _ := cc.cs.GetAllClients()
	if len(clients) == 0 {
		common.JSONResponse(w, http.StatusOK, common.ApiResponse{
			Message: "CLIENTS_FETCHED_SUCCESSFULLY",
			Data:    clients,
		})
		return
	}
	common.JSONResponse(w, http.StatusOK, common.ApiResponse{
		Message: "CLIENTS_FETCHED_SUCCESSFULLY",
		Data:    clients,
	})
}

func (cc *ClientController) GetClientById(w http.ResponseWriter, r *http.Request) {
	requestID := mux.Vars(r)["clientID"]
	clientID, _ := uuid.Parse(requestID)
	client, err := cc.cs.GetClientById(clientID)

	if err != nil {
		common.JSONResponse(w, http.StatusOK, common.ApiResponse{
			Error: &common.ErrorResponse{
				Message: "ERROR_FETCHING_CLIENT",
			},
		})
		return
	}

	if (client == models.GetClientRequest{}) {
		common.JSONResponse(w, http.StatusNotFound, common.ApiResponse{
			Message: "CLIENT_NOT_FOUND",
			Data:    client,
		})
		return
	}

	common.JSONResponse(w, http.StatusNotFound, common.ApiResponse{
		Message: "CLIENT_FOUND",
		Data:    client,
	})
}

func (cc *ClientController) UpdateClient(w http.ResponseWriter, r *http.Request) {
	requestID := mux.Vars(r)["clientID"]
	clientID, err := uuid.Parse(requestID)
	if err != nil {
		common.JSONResponse(w, http.StatusBadRequest, common.ApiResponse{
			Error: &common.ErrorResponse{
				Message: "Invalid client ID",
				Details: err.Error(),
			},
		})
		return
	}

	var request models.UpdateClientRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		common.JSONResponse(w, http.StatusBadRequest, common.ApiResponse{
			Error: &common.ErrorResponse{
				Message: "Invalid request body",
				Details: err.Error(),
			},
		})
		return
	}

	if errors := utils.ValidateEntity(request); errors != nil {
		common.JSONResponse(w, http.StatusUnprocessableEntity, common.ApiResponse{
			Error: &common.ErrorResponse{
				Message: "Invalid data",
				Details: errors,
			},
		})
		return
	}

	updatedClient := models.NewClient(
		request.DocumentType,
		request.Document,
		request.Name,
		request.Surname,
		request.Email,
		request.PhoneNumber,
		request.Address,
	)
	updatedClient.Person.ID = clientID

	err = cc.cs.UpdateClient(clientID, *updatedClient)
	if err != nil {
		common.JSONResponse(w, http.StatusNotFound, common.ApiResponse{
			Error: &common.ErrorResponse{
				Message: "CLIENT_NOT_FOUND",
			},
		})
		return
	}

	common.JSONResponse(w, http.StatusOK, common.ApiResponse{
		Message: "CLIENT_UPDATED_SUCCESSFULLY",
		Data:    request,
	})
}

func (cc *ClientController) DeleteClient(w http.ResponseWriter, r *http.Request) {
	requestID := mux.Vars(r)["clientID"]
	clientID, err := uuid.Parse(requestID)
	if err != nil {
		common.JSONResponse(w, http.StatusBadRequest, common.ApiResponse{
			Error: &common.ErrorResponse{
				Message: "Invalid client ID",
				Details: err.Error(),
			},
		})
		return
	}

	err = cc.cs.DeleteClient(clientID)
	if err != nil {
		common.JSONResponse(w, http.StatusNotFound, common.ApiResponse{
			Error: &common.ErrorResponse{
				Message: "CLIENT_NOT_FOUND",
			},
		})
		return
	}

	common.JSONResponse(w, http.StatusOK, common.ApiResponse{
		Message: "CLIENT_DELETED_SUCCESSFULLY",
	})
}
