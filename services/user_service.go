package services

import (
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/models"
	"AuthInGo/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById(id string) (*models.User, error)
	CreateUser(payload *dto.CreateUserRequestDTO) (*models.User, error)
	LoginUser(payload *dto.LoginUserRequestDTO) (string, error)
}

type UserServiceImp struct {
	userRepository db.UsersRepository
}

func NewUserService(_userRepository db.UsersRepository) UserService {
	return &UserServiceImp{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImp) GetUserById(id string) (*models.User, error) {
	fmt.Println("Fetching user in service layer")
	user, err := u.userRepository.GetById(id)
	if err != nil {
		fmt.Println("Error fetching user in service layer:", err)						
		return nil, err
	}

	fmt.Println("User fetched in service layer:", user)
	return user, nil
}

func (u *UserServiceImp) CreateUser(payload *dto.CreateUserRequestDTO) (*models.User, error) {
	fmt.Println("Creating user in service layer")

	hashedpassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		fmt.Println("Error hashing password in service layer:", err)
		return nil, err
	}
	user, err := u.userRepository.Create(
		payload.Username,
		payload.Email,
		hashedpassword,
	)

	if err != nil {
		fmt.Println("Error creating user in service layer:", err)
		return nil, err
	}

	fmt.Println("User created in service layer:", user)
	return user, nil
}

func (u *UserServiceImp) LoginUser(payload *dto.LoginUserRequestDTO) (string, error) {
	email := payload.Email
	password := payload.Password
	user, err := u.userRepository.GetByEmail(email)

	if err != nil {
		fmt.Println("Error fetching user by email in service layer:", err)
		return "", err
	}

	if user == nil {
		fmt.Println("User not found with email:", email)
		return "", nil
	}
	
	isPasswordValid := utils.CheckPasswordHash(password, user.Password)

	if !isPasswordValid {
		fmt.Println("Invalid password for user with email:", email)
		return "", nil
	}

	jwtPayload := jwt.MapClaims{
		"email": user.Email,
		"id": user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)

	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "default_secret")))

	if err != nil {
		fmt.Println("Error generating JWT token:", err)
		return "", err
	}

	fmt.Println("Generated JWT token:", tokenString)

	return tokenString, nil
}


