// @title Microservicio de usuario UDC
// @version 1.0
// @description Micro para gestión de usuarios, clientes y técnicos
// @host localhost:8080
// @BasePath /user-service/v1
package main

import (
	"context"
	"log"
	"net/http"
	_ "user-service-ucd/docs"
	"user-service-ucd/src/app"
	"user-service-ucd/src/app/controllers"
	"user-service-ucd/src/app/repository"
	"user-service-ucd/src/app/services"
	"user-service-ucd/src/database"

	"go.uber.org/fx"
)

func RunServer(lc fx.Lifecycle, routes *app.Routes) {
	server := &http.Server{
		Addr:    ":8080",
		Handler: routes.Router,
	}

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			log.Println("Server listening on :8080")
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(_ context.Context) error {
			log.Println("Server shutting down...")
			return server.Close()
		},
	})
}

func main() {

	// Inyección de dependencias
	app := fx.New(
		fx.Provide(
			// Inicializa la conexión con la BD
			database.InitDB,
			
			//Dependencias clientes
			repository.NewClientRepository,
			services.NewClientService,
			controllers.NewClientController,

			//Dependencias Técnicos
			repository.NewTechnicianRepository,
			services.NewTechnicianService,
			controllers.NewTechnicianController,

			// Dependencias auth
			services.NewAuthService,
			controllers.NewAuthController,
			
			app.NewRouter,
		),
		fx.Invoke(RunServer),
	)

	app.Run()
}
