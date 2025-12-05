package services

import (
	db "AuthInGo/db/repositories"
	"fmt"
)

type UserService interface {
	CreateUser() error
}

type UserServiceImp struct {
	userRepository db.UsersRepository
}

func (u *UserServiceImp) CreateUser() error {
	fmt.Println("Creating user in service layer")
	u.userRepository.Create()
	return nil
}

func NewUserService(_userRepository db.UsersRepository) UserService {
	return &UserServiceImp{
		userRepository: _userRepository,
	}
}