package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"
	"AuthInGo/utils"
	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(r chi.Router)
}

func SetupRoutes(UserRouter Router, RoleRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()

	chiRouter.Use(middlewares.RequestLogger)

	chiRouter.Use(middlewares.RateLimiterMiddleware)

	chiRouter.Get("/ping", controllers.PingHandler) 

	chiRouter.HandleFunc("/fakestoreservice/*", utils.ProxyToService("https://fakestoreapi.com/", "/fakestoreservice"))

	UserRouter.Register(chiRouter)
	RoleRouter.Register(chiRouter)

	return chiRouter
} 	
