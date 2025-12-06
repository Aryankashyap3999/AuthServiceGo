package services

import (
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById() error
	CreateUser() error
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

func (u *UserServiceImp) GetUserById() error {
	fmt.Println("Fetching user in service layer")
	u.userRepository.GetById()
	return nil
}

func (u *UserServiceImp) CreateUser() error {
	fmt.Println("Creating user in service layer")
	password := "plain_password" // In real scenario, hash the password properly
	hassedPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Println("Error hashing password in service layer:", err)
		return err
	}
	u.userRepository.Create(
		"username_example_3",
		"user3@example.com",
		hassedPassword,
	)
	return nil
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
		"email": email,
		"password": password,
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


