package controllers

import (
	"encoding/json"
	"net/http"
	"user-service-ucd/src/app/models"
	"user-service-ucd/src/app/services"
	"user-service-ucd/src/common"
	"user-service-ucd/utils"

	"github.com/gorilla/mux"
)

type TechnicianController struct {
	ts *services.TechnicianService
}

func NewTechnicianController(ts *services.TechnicianService) *TechnicianController {
	return &TechnicianController{
		ts: ts,
	}
}

// Crear técnico
func (tc *TechnicianController) AddNewTechnician(w http.ResponseWriter, r *http.Request) {
	var createTechRequest models.CreateTechnicianRequest
	if err := json.NewDecoder(r.Body).Decode(&createTechRequest); err != nil {
		common.JSONResponse(w, http.StatusBadRequest, common.ApiResponse{
			Error: &common.ErrorResponse{
				Message: "Invalid request body",
				Details: err.Error(),
			},
		})
		return
	}

	if errors := utils.ValidateEntity(createTechRequest); errors != nil {
		common.JSONResponse(w, http.StatusUnprocessableEntity, common.ApiResponse{
			Error: &common.ErrorResponse{
				Message: "Invalid data",
				Details: errors,
			},
		})
		return
	}

	technician, err := models.NewTechnician(createTechRequest)
	if err != nil {
		common.JSONResponse(w, http.StatusInternalServerError, common.ApiResponse{
			Error: &common.ErrorResponse{
				Message: "Error creating technician",
				Details: err.Error(),
			},
		})
		return
	}

	tc.ts.NewTechnician(*technician)
	common.JSONResponse(w, http.StatusCreated, common.ApiResponse{
		Message: "TECHNICIAN_CREATED_SUCCESSFULLY",
		Data:    createTechRequest,
	})
}

// Obtener todos los técnicos
func (tc *TechnicianController) GetAllTechnicians(w http.ResponseWriter, r *http.Request) {
	technicians, _ := tc.ts.GetAllTechnicians()
	common.JSONResponse(w, http.StatusOK, common.ApiResponse{
		Message: "TECHNICIANS_FETCHED_SUCCESSFULLY",
		Data:    technicians,
	})
}

// Obtener técnico por ID
func (tc *TechnicianController) GetTechnicianById(w http.ResponseWriter, r *http.Request) {
	requestID := mux.Vars(r)["technicianID"]
	technician, err := tc.ts.GetTechnicianById(requestID)

	if err != nil {
		common.JSONResponse(w, http.StatusNotFound, common.ApiResponse{
			Error: &common.ErrorResponse{
				Message: "TECHNICIAN_NOT_FOUND",
			},
		})
		return
	}

	common.JSONResponse(w, http.StatusOK, common.ApiResponse{
		Message: "TECHNICIAN_FOUND",
		Data:    technician,
	})
}

// Actualizar técnico
func (tc *TechnicianController) UpdateTechnician(w http.ResponseWriter, r *http.Request) {
	technicianId := mux.Vars(r)["technicianID"]

	var request models.UpdateTechnicianRequest
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

	updatedTechnician, err := models.NewTechnicianUpdated(request, technicianId)
	if err != nil {
		common.JSONResponse(w, http.StatusInternalServerError, common.ApiResponse{
			Error: &common.ErrorResponse{
				Message: "Error updating technician",
				Details: err.Error(),
			},
		})
		return
	}

	err = tc.ts.UpdateTechnician(technicianId, *updatedTechnician)
	if err != nil {
		common.JSONResponse(w, http.StatusNotFound, common.ApiResponse{
			Error: &common.ErrorResponse{
				Message: "TECHNICIAN_NOT_FOUND",
			},
		})
		return
	}

	common.JSONResponse(w, http.StatusOK, common.ApiResponse{
		Message: "TECHNICIAN_UPDATED_SUCCESSFULLY",
		Data:    request,
	})
}

// Eliminar técnico
func (tc *TechnicianController) DeleteTechnician(w http.ResponseWriter, r *http.Request) {
	technicianID := mux.Vars(r)["technicianID"]

	err := tc.ts.DeleteTechnician(technicianID)
	if err != nil {
		common.JSONResponse(w, http.StatusNotFound, common.ApiResponse{
			Error: &common.ErrorResponse{
				Message: "TECHNICIAN_NOT_FOUND",
			},
		})
		return
	}

	common.JSONResponse(w, http.StatusOK, common.ApiResponse{
		Message: "TECHNICIAN_DELETED_SUCCESSFULLY",
	})
}

// Cambiar contraseña
func (tc *TechnicianController) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var request models.ChangePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		common.JSONResponse(w, http.StatusBadRequest, common.ApiResponse{
			Error: &common.ErrorResponse{
				Message: "Invalid request body",
				Details: err.Error(),
			},
		})
		return
	}

	if request.Username == "" || request.NewPassword == "" {
		common.JSONResponse(w, http.StatusBadRequest, common.ApiResponse{
			Error: &common.ErrorResponse{
				Message: "Username and new password are required",
			},
		})
		return
	}

	err := tc.ts.ChangePassword(request.Username, request.NewPassword)
	if err != nil {
		common.JSONResponse(w, http.StatusNotFound, common.ApiResponse{
			Error: &common.ErrorResponse{
				Message: "TECHNICIAN_NOT_FOUND",
				Details: err.Error(),
			},
		})
		return
	}

	common.JSONResponse(w, http.StatusOK, common.ApiResponse{
		Message: "PASSWORD_CHANGED_SUCCESSFULLY",
	})
}
