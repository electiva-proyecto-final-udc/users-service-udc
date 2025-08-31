package main

import (
	"context"
	"log"
	"net/http"
	"user-service-ucd/src/app"
	"user-service-ucd/src/app/controllers"
	"user-service-ucd/src/app/repository"
	"user-service-ucd/src/app/services"

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
			
			//Dependencias clientes
			repository.NewClientRepository,
			services.NewClientService,
			controllers.NewClientController,

			app.NewRouter,
		),
		fx.Invoke(RunServer),
	)

	app.Run()
}
