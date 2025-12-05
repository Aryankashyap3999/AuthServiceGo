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

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUserById called in controller")
	uc.UserService.GetUserById()
	w.Write([]byte("User fetched successfully"))
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUserById called in controller")
	uc.UserService.CreateUser()
	w.Write([]byte("User Created successfully"))
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginUser called in controller")
	uc.UserService.LoginUser()
	w.Write([]byte("User Logged in successfully"))
}