package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/services"
	"AuthInGo/utils"
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

	w.Write([]byte("User fetched successfully"))
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	payload := r.Context().Value("create_payload").(dto.CreateUserRequestDTO)

	fmt.Println("Payload received in controller:", payload)

	user, err := uc.UserService.CreateUser(&payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error creating user: ", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusCreated, "User created successfully", user)
	fmt.Println("User created in controller:", user)
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginUser called in controller")

	payload := r.Context().Value("login_payload").(dto.LoginUserRequestDTO)

	fmt.Println("Payload in controller:", payload)

	if validationErr := utils.Validator.Struct(payload); validationErr != nil {
		fmt.Println("Validation error in controller:", validationErr)
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation error: ", validationErr)
		return
	}

	jwtToken, err := uc.UserService.LoginUser(&payload)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusUnauthorized, "Login failed: ", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User logged in successfully", jwtToken)
}
