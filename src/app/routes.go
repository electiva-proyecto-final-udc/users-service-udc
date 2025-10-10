package app

import (
	"net/http"
	"user-service-ucd/src/app/controllers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Routes struct {
	Router http.Handler // 👈 cambiamos a http.Handler, no *mux.Router
}

func NewRouter(
	clientController *controllers.ClientController,
	technicianController *controllers.TechnicianController,
	authController *controllers.AuthController,
) *Routes {

	r := mux.NewRouter()
	api := r.PathPrefix("/user-service/v1").Subrouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

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

	// Configurar CORS globalmente
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // React
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Aplicar CORS al router
	handler := c.Handler(r)

	return &Routes{
		Router: handler, // 👈 ahora Router es un http.Handler
	}
}
