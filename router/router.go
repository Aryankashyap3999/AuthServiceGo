package router

import "github.com/go-chi/chi/v5"
import "AuthInGo/controllers"

type Router interface {
	Register(r chi.Router)
}

func SetupRoutes(UserRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()

	chiRouter.Get("/ping", controllers.PingHandler) 

	UserRouter.Register(chiRouter)

	return chiRouter
} 	