package app

import (
	"net/http"
	"user-service-ucd/src/app/controllers"
	"user-service-ucd/src/middleware"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Routes struct {
	Router http.Handler 
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

	// Subrouter protegido
	protected := api.NewRoute().Subrouter()
	protected.Use(middleware.VerifyToken)

	// Rutas Cliente
	protected.HandleFunc("/clients", clientController.GetAllClients).Methods("GET")
	protected.HandleFunc("/clients/{clientID}", clientController.GetClientById).Methods("GET")
	protected.HandleFunc("/createClient", clientController.AddNewClient).Methods("POST")
	protected.HandleFunc("/updateClient/{clientID}", clientController.UpdateClient).Methods("PATCH")
	protected.HandleFunc("/deleteClient/{clientID}", clientController.DeleteClient).Methods("DELETE")

	// Rutas Técnico
	protected.HandleFunc("/technicians", technicianController.GetAllTechnicians).Methods("GET")
	protected.HandleFunc("/technician/{technicianID}", technicianController.GetTechnicianById).Methods("GET")
	protected.HandleFunc("/createTechnician", technicianController.AddNewTechnician).Methods("POST")
	protected.HandleFunc("/updateTechnician/{technicianID}", technicianController.UpdateTechnician).Methods("PATCH")
	protected.HandleFunc("/deleteTechnician/{technicianID}", technicianController.DeleteTechnician).Methods("DELETE")
	protected.HandleFunc("/changePassword", technicianController.ChangePassword).Methods("PATCH")

	// Configurar CORS globalmente
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Front
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Aplicar CORS al router
	handler := c.Handler(r)

	return &Routes{
		Router: handler,
	}
}
