package controllers

import (
	"AuthInGo/services"
	"fmt"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}	

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	uc.UserService.CreateUser()
	fmt.Println("User registration handled in controller")
	w.Write([]byte("User registered"))
}