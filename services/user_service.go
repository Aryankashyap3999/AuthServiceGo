package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"fmt"
)

type UserService interface {
	GetUserById() error
	CreateUser() error
	LoginUser() error
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

func (u *UserServiceImp) LoginUser() error {
	response := utils.CheckPasswordHash("plain_password_wrong", "$2a$10$l/kfXeiblOcYPwTjJ.DQFOfEagz.t2JCdDTXTFT8LsK.baVa/p2zG")
	fmt.Println("Password match status:", response)
	return nil
}


