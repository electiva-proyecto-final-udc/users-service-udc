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
) *Routes {

	r := mux.NewRouter()
	api := r.PathPrefix("/user-service/v1").Subrouter()

	// Rutas Cliente
	api.HandleFunc("/clients", clientController.GetAllClients).Methods("GET")
	api.HandleFunc("/clients", clientController.AddNewClient).Methods("POST")

	return &Routes{
		Router: r,
	}
}
