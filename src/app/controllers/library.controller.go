package controllers

import (
	"net/http"
	"user-service-ucd/src/app/services"
	"user-service-ucd/src/common"
)

type LibraryController struct {
	ls *services.LibraryService
}

func NewLibraryController(ls *services.LibraryService) *LibraryController {
	return &LibraryController{
		ls: ls,
	}
}

// GetDocumenTypes obtiene todos los tipos de documentos
// @Summary      Listar Tipos de documentos
// @Description  Obtiene la lista completa de los tipos de documentos
// @Tags         technicians
// @Produce      json
// @Success      200 {object} common.ApiResponse{data=[]models.LibraryModel}
// @Router       /library/document_types [get]

func (lc *LibraryController) GetDocumentTypes(w http.ResponseWriter, r *http.Request) {
	documentTypes, err := lc.ls.GetDocumentTypes()
	if err != nil {
		common.JSONResponse(w, http.StatusNotFound, common.ApiResponse{
			Error: &common.ErrorResponse{
				Code:    400,
				Message: "ERROR GETTING DOCUMENT TYPE",
				Details: err.Error(),
			},
		})
		return
	}

	common.JSONResponse(w, http.StatusOK, common.ApiResponse{
		Message: "DOCUMENT_TYPES_FETCHED_SUCCESSFULLY",
		Data:    documentTypes,
	})
}
