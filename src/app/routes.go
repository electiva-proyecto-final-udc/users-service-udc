package app

import (
	"user-service-ucd/src/app/controllers"

	"github.com/gorilla/mux"
)

type Routes struct {
	Router *mux.Router
}

func NewRouter(
	clientController *controllers.ClientController,
	technicianController *controllers.TechnicianController,
	authController *controllers.AuthController,
) *Routes {

	r := mux.NewRouter()
	api := r.PathPrefix("/user-service/v1").Subrouter()

	// Auth
	api.HandleFunc("/auth/login", authController.Login).Methods("POST")

	// Rutas Cliente
	api.HandleFunc("/clients", clientController.GetAllClients).Methods("GET")
	api.HandleFunc("/clients/{clientID}", clientController.GetClientById).Methods("GET")
	api.HandleFunc("/createClient", clientController.AddNewClient).Methods("POST")
	api.HandleFunc("/updateClient/{clientID}", clientController.UpdateClient).Methods("PUT")
	api.HandleFunc("/deleteClient/{clientID}", clientController.DeleteClient).Methods("DELETE")

	// Rutas Técnico
	api.HandleFunc("/technicians", technicianController.GetAllTechnicians).Methods("GET")
	api.HandleFunc("/technician/{technicianID}", technicianController.GetTechnicianById).Methods("GET")
	api.HandleFunc("/createTechnician", technicianController.AddNewTechnician).Methods("POST")
	api.HandleFunc("/updateTechnician/{technicianID}", technicianController.UpdateTechnician).Methods("PUT")
	api.HandleFunc("/deleteTechnician/{technicianID}", technicianController.DeleteTechnician).Methods("DELETE")
	api.HandleFunc("/changePassword", technicianController.ChangePassword).Methods("PATCH")

	return &Routes{
		Router: r,
	}
}
