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
	api.HandleFunc("/clients/{clientID}", clientController.GetClientById).Methods("GET")
	api.HandleFunc("/createClient", clientController.AddNewClient).Methods("POST")
	api.HandleFunc("/updateClient/{clientID}", clientController.UpdateClient).Methods("PUT")
	api.HandleFunc("/deleteClient/{clientID}", clientController.DeleteClient).Methods("DELETE")

	return &Routes{
		Router: r,
	}
}
