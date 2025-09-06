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

// AddNewTechnician crea un nuevo técnico
// @Summary      Crear técnico
// @Description  Crea un nuevo técnico en el sistema
// @Tags         technicians
// @Accept       json
// @Produce      json
// @Param        request body models.CreateTechnicianRequest true "Datos del técnico"
// @Success      201 {object} common.ApiResponse{data=models.Technician}
// @Failure      400 {object} common.ApiResponse{error=common.ErrorResponse}
// @Failure      422 {object} common.ApiResponse{error=common.ErrorResponse}
// @Failure      500 {object} common.ApiResponse{error=common.ErrorResponse}
// @Router       /createTechnician [post]
func (tc *TechnicianController) AddNewTechnician(w http.ResponseWriter, r *http.Request) {
	var createTechRequest models.CreateTechnicianRequest
	if err := json.NewDecoder(r.Body).Decode(&createTechRequest); err != nil {
		common.JSONResponse(w, http.StatusBadRequest, common.ApiResponse{
			Error: &common.ErrorResponse{
				Code: 400,
				Message: "Invalid request body",
				Details: err.Error(),
			},
		})
		return
	}

	if errors := utils.ValidateEntity(createTechRequest); errors != nil {
		common.JSONResponse(w, http.StatusUnprocessableEntity, common.ApiResponse{
			Error: &common.ErrorResponse{
				Code: 422,
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
				Code: 500,
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

// GetAllTechnicians obtiene todos los técnicos
// @Summary      Listar técnicos
// @Description  Obtiene la lista completa de técnicos
// @Tags         technicians
// @Produce      json
// @Success      200 {object} common.ApiResponse{data=[]models.Technician}
// @Router       /technicians [get]
func (tc *TechnicianController) GetAllTechnicians(w http.ResponseWriter, r *http.Request) {
	technicians, _ := tc.ts.GetAllTechnicians()
	common.JSONResponse(w, http.StatusOK, common.ApiResponse{
		Message: "TECHNICIANS_FETCHED_SUCCESSFULLY",
		Data:    technicians,
	})
}

// GetTechnicianById obtiene un técnico por ID
// @Summary      Obtener técnico
// @Description  Busca un técnico por su ID
// @Tags         technicians
// @Produce      json
// @Param        technicianID path string true "ID del técnico" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"
// @Success      200 {object} common.ApiResponse{data=models.Technician}
// @Failure      404 {object} common.ApiResponse{error=common.ErrorResponse}
// @Router       /technician/{technicianID} [get]
func (tc *TechnicianController) GetTechnicianById(w http.ResponseWriter, r *http.Request) {
	requestID := mux.Vars(r)["technicianID"]
	technician, err := tc.ts.GetTechnicianById(requestID)

	if err != nil {
		common.JSONResponse(w, http.StatusNotFound, common.ApiResponse{
			Error: &common.ErrorResponse{
				Code: 404,
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

// UpdateTechnician actualiza un técnico
// @Summary      Actualizar técnico
// @Description  Actualiza los datos de un técnico existente
// @Tags         technicians
// @Accept       json
// @Produce      json
// @Param        technicianID path string true "ID del técnico" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"
// @Param        request body models.UpdateTechnicianRequest true "Datos del técnico"
// @Success      200 {object} common.ApiResponse{data=models.Technician}
// @Failure      400 {object} common.ApiResponse{error=common.ErrorResponse}
// @Failure      404 {object} common.ApiResponse{error=common.ErrorResponse}
// @Failure      422 {object} common.ApiResponse{error=common.ErrorResponse}
// @Failure      500 {object} common.ApiResponse{error=common.ErrorResponse}
// @Router       /updateTechnician/{technicianID} [put]
func (tc *TechnicianController) UpdateTechnician(w http.ResponseWriter, r *http.Request) {
	technicianId := mux.Vars(r)["technicianID"]

	var request models.UpdateTechnicianRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		common.JSONResponse(w, http.StatusBadRequest, common.ApiResponse{
			Error: &common.ErrorResponse{
				Code: 400,
				Message: "Invalid request body",
				Details: err.Error(),
			},
		})
		return
	}

	if errors := utils.ValidateEntity(request); errors != nil {
		common.JSONResponse(w, http.StatusUnprocessableEntity, common.ApiResponse{
			Error: &common.ErrorResponse{
				Code: 422,
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
				Code: 500,
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
				Code: 404,
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

// DeleteTechnician elimina un técnico
// @Summary      Eliminar técnico
// @Description  Elimina un técnico por su ID
// @Tags         technicians
// @Produce      json
// @Param        technicianID path string true "ID del técnico" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"
// @Success      200 {object} common.ApiResponse "Ejemplo: {\"message\":\"TECHNICIAN_DELETED_SUCCESSFULLY\"}"
// @Failure      404 {object} common.ApiResponse{error=common.ErrorResponse}
// @Router       /deleteTechnician/{technicianID} [delete]
func (tc *TechnicianController) DeleteTechnician(w http.ResponseWriter, r *http.Request) {
	technicianID := mux.Vars(r)["technicianID"]

	err := tc.ts.DeleteTechnician(technicianID)
	if err != nil {
		common.JSONResponse(w, http.StatusNotFound, common.ApiResponse{
			Error: &common.ErrorResponse{
				Code: 404,
				Message: "TECHNICIAN_NOT_FOUND",
			},
		})
		return
	}

	common.JSONResponse(w, http.StatusOK, common.ApiResponse{
		Message: "TECHNICIAN_DELETED_SUCCESSFULLY",
	})
}

// ChangePassword cambia la contraseña de un técnico
// @Summary      Cambiar contraseña
// @Description  Cambia la contraseña de un técnico existente
// @Tags         technicians
// @Accept       json
// @Produce      json
// @Param        request body models.ChangePasswordRequest true "Datos para cambiar contraseña"
// @Success      200 {object} common.ApiResponse "Ejemplo: {\"message\":\"PASSWORD_CHANGED_SUCCESSFULLY\"}"
// @Failure      400 {object} common.ApiResponse{error=common.ErrorResponse}
// @Failure      404 {object} common.ApiResponse{error=common.ErrorResponse}
// @Router       /changePassword [patch]
func (tc *TechnicianController) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var request models.ChangePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		common.JSONResponse(w, http.StatusBadRequest, common.ApiResponse{
			Error: &common.ErrorResponse{
				Code: 400,
				Message: "Invalid request body",
				Details: err.Error(),
			},
		})
		return
	}

	if request.Username == "" || request.NewPassword == "" {
		common.JSONResponse(w, http.StatusBadRequest, common.ApiResponse{
			Error: &common.ErrorResponse{
				Code: 400,
				Message: "Username and new password are required",
			},
		})
		return
	}

	err := tc.ts.ChangePassword(request.Username, request.NewPassword)
	if err != nil {
		common.JSONResponse(w, http.StatusNotFound, common.ApiResponse{
			Error: &common.ErrorResponse{
				Code: 404,
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
