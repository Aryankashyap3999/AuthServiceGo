package app

import (
	config "AuthInGo/config/env"
	"AuthInGo/controllers"
	db "AuthInGo/db/repositories"
	"AuthInGo/router"
	"AuthInGo/services"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string
}

type Application struct {
	Config Config
	Store db.Storage
}	

func NewConfig() Config {

	port := config.GetString("port", "8080")

	return Config{
		Addr: port,
	}	
}

func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
		Store: *db.NewStorage(),
	}
}


func (app *Application) Run() error {

	ur := db.NewUserRepository()
	us := services.NewUserService(ur)
	uc := controllers.NewUserController(us)
	uRouter := router.NewUserRouter(uc)

	server := &http.Server{
		Addr:         app.Config.Addr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router.SetupRoutes(uRouter),
	}

	fmt.Println("Starting server on", app.Config.Addr)

	return server.ListenAndServe()
}

