package controllers

import (
	"encoding/json"
	"net/http"
	"user-service-ucd/src/app/dto"
	"user-service-ucd/src/app/models"
	"user-service-ucd/src/app/services"
	"user-service-ucd/src/common"
	"user-service-ucd/utils"

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

// AddNewClient crea un nuevo cliente
// @Summary      Crear cliente
// @Description  Crea un nuevo cliente en el sistema
// @Tags         clients
// @Accept       json
// @Produce      json
// @Param        request body dto.CreateClientRequest true "Datos del cliente"
// @Success      201 {object} common.ApiResponse{data=models.Client}
// @Failure      400 {object} common.ApiResponse{error=common.ErrorResponse}
// @Failure      422 {object} common.ApiResponse{error=common.ErrorResponse}
// @Router       /createClient [post]
func (cc *ClientController) AddNewClient(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateClientRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		common.JSONResponse(w, http.StatusBadRequest, common.ApiResponse{
			Error: &common.ErrorResponse{
				Code:    400,
				Message: "Invalid request body",
				Details: err.Error(),
			},
		})
		return
	}

	if errors := utils.ValidateEntity(request); errors != nil {
		common.JSONResponse(w, http.StatusUnprocessableEntity, common.ApiResponse{
			Error: &common.ErrorResponse{
				Code:    422,
				Message: "Invalid data",
				Details: errors,
			},
		})
		return
	}

	id, err := cc.cs.NewClient(request)
	if  err != nil {
		common.JSONResponse(w, http.StatusBadRequest, common.ApiResponse{
			Error: &common.ErrorResponse{
				Code:    400,
				Message: "ERROR CREATING CLIENT",
				Details: err.Error(),
			},
		})
		return
	}
	request.ID = id
	common.JSONResponse(w, http.StatusCreated, common.ApiResponse{
		Message: "CLIENT_CREATED_SUCCESSFULLY",
		Data:    request,
	})
}

// GetAllClients obtiene todos los clientes
// @Summary      Listar clientes
// @Description  Obtiene la lista completa de clientes
// @Tags         clients
// @Produce      json
// @Success      200 {object} common.ApiResponse{data=[]models.Client}
// @Router       /clients [get]
func (cc *ClientController) GetAllClients(w http.ResponseWriter, r *http.Request) {
	clients, _ := cc.cs.GetAllClients()
	if len(clients) == 0 {
		common.JSONResponse(w, http.StatusOK, common.ApiResponse{
			Message: "NO_CLIENTS_FOUND",
			Data:    clients,
		})
		return
	}
	common.JSONResponse(w, http.StatusOK, common.ApiResponse{
		Message: "CLIENTS_FETCHED_SUCCESSFULLY",
		Data:    clients,
	})
}

// GetClientById obtiene un cliente por ID
// @Summary      Obtener cliente
// @Description  Busca un cliente por su UUID
// @Tags         clients
// @Produce      json
// @Param        clientID path string true "ID del cliente" example:"550e8400-e29b-41d4-a716-446655440000"
// @Success      200 {object} common.ApiResponse{data=models.Client}
// @Failure      404 {object} common.ApiResponse{error=common.ErrorResponse}
// @Failure      500 {object} common.ApiResponse{error=common.ErrorResponse}
// @Router       /clients/{clientID} [get]
func (cc *ClientController) GetClientById(w http.ResponseWriter, r *http.Request) {
	requestID := mux.Vars(r)["clientID"]
	client, err := cc.cs.GetClientById(requestID)

	if err != nil {
		common.JSONResponse(w, http.StatusInternalServerError, common.ApiResponse{
			Error: &common.ErrorResponse{
				Code:    500,
				Message: "ERROR_FETCHING_CLIENT",
				Details: err.Error(),
			},
		})
		return
	}

	if (client == models.ClientDataView{}) {
		common.JSONResponse(w, http.StatusNotFound, common.ApiResponse{
			Error: &common.ErrorResponse{
				Code:    404,
				Message: "CLIENT_NOT_FOUND",
			},
		})
		return
	}

	common.JSONResponse(w, http.StatusOK, common.ApiResponse{
		Message: "CLIENT_FETCHED_SUCCESFULLY",
		Data:    client,
	})
}

// UpdateClient actualiza un cliente
// @Summary      Actualizar cliente
// @Description  Actualiza los datos de un cliente existente
// @Tags         clients
// @Accept       json
// @Produce      json
// @Param        clientID path string true "ID del cliente" example:"550e8400-e29b-41d4-a716-446655440000"
// @Param        request body dto.UpdateClientRequest true "Datos del cliente"
// @Success      200 {object} common.ApiResponse{data=models.Client}
// @Failure      400 {object} common.ApiResponse{error=common.ErrorResponse}
// @Failure      404 {object} common.ApiResponse{error=common.ErrorResponse}
// @Failure      422 {object} common.ApiResponse{error=common.ErrorResponse}
// @Router       /updateClient/{clientID} [put]
func (cc *ClientController) UpdateClient(w http.ResponseWriter, r *http.Request) {
	requestID := mux.Vars(r)["clientID"]

	if requestID == "" {
		common.JSONResponse(w, http.StatusBadRequest, common.ApiResponse{
			Error: &common.ErrorResponse{
				Code:    400,
				Message: "Invalid request",
			},
		})
		return
	}

	var updatedClient dto.UpdateClientRequest
	if err := json.NewDecoder(r.Body).Decode(&updatedClient); err != nil {
		common.JSONResponse(w, http.StatusBadRequest, common.ApiResponse{
			Error: &common.ErrorResponse{
				Code:    400,
				Message: "Invalid request body",
				Details: err.Error(),
			},
		})
		return
	}

	if errors := utils.ValidateEntity(updatedClient); errors != nil {
		common.JSONResponse(w, http.StatusUnprocessableEntity, common.ApiResponse{
			Error: &common.ErrorResponse{
				Code:    422,
				Message: "Invalid data",
				Details: errors,
			},
		})
		return
	}

	err := cc.cs.UpdateClient(requestID, updatedClient)
	if err != nil {
		common.JSONResponse(w, http.StatusNotFound, common.ApiResponse{
			Error: &common.ErrorResponse{
				Code:    404,
				Message: "CLIENT_NOT_FOUND",
				Details: err.Error(),
			},
		})
		return
	}

	common.JSONResponse(w, http.StatusOK, common.ApiResponse{
		Message: "CLIENT_UPDATED_SUCCESSFULLY",
		Data:    updatedClient,
	})
}

// DeleteClient elimina un cliente
// @Summary      Eliminar cliente
// @Description  Elimina un cliente por su UUID
// @Tags         clients
// @Produce      json
// @Param        clientID path string true "ID del cliente" example:"550e8400-e29b-41d4-a716-446655440000"
// @Success      200 {object} common.ApiResponse "Ejemplo: {\"message\":\"CLIENT_DELETED_SUCCESSFULLY\"}"
// @Failure      400 {object} common.ApiResponse{error=common.ErrorResponse}
// @Failure      404 {object} common.ApiResponse{error=common.ErrorResponse}
// @Router       /deleteClient/{clientID} [delete]
func (cc *ClientController) DeleteClient(w http.ResponseWriter, r *http.Request) {
	clientID := mux.Vars(r)["clientID"]

	err := cc.cs.DeleteClient(clientID)

	if err != nil {
		common.JSONResponse(w, http.StatusNotFound, common.ApiResponse{
			Error: &common.ErrorResponse{
				Code:    404,
				Message: "CLIENT_NOT_FOUND",
			},
		})
		return
	}

	common.JSONResponse(w, http.StatusOK, common.ApiResponse{
		Message: "CLIENT_DELETED_SUCCESSFULLY",
	})
}
