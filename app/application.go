package app

import (
	config "AuthInGo/config/env"
	"AuthInGo/controllers"
	repo "AuthInGo/db/repositories"
	"AuthInGo/router"
	"AuthInGo/services"
	"fmt"
	"net/http"
	"time"
	dbConfig "AuthInGo/config/db"
)

type Config struct {
	Addr string
}

type Application struct {
	Config Config
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
	}
}


func (app *Application) Run() error {

	db, err := dbConfig.SetupDB()

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return err
	}

	ur := repo.NewUserRepository(db)
	rr := repo.NewRoleRepository(db)
	us := services.NewUserService(ur)
	rs := services.NewRoleService(rr)
	uc := controllers.NewUserController(us)
	rc := controllers.NewRoleController(rs)
	uRouter := router.NewUserRouter(uc)
	rRouter := router.NewRoleRouter(rc)

	server := &http.Server{
		Addr:         app.Config.Addr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router.SetupRoutes(uRouter, rRouter),
	}

	fmt.Println("Starting server on", app.Config.Addr)

	return server.ListenAndServe()
}

