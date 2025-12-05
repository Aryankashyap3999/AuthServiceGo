package services

import (
	db "AuthInGo/db/repositories"
	"fmt"
)

type UserService interface {
	GetUserById() error
}

type UserServiceImp struct {
	userRepository db.UsersRepository
}

func (u *UserServiceImp) GetUserById() error {
	fmt.Println("Fetching user in service layer")
	u.userRepository.GetById()
	return nil
}

func NewUserService(_userRepository db.UsersRepository) UserService {
	return &UserServiceImp{
		userRepository: _userRepository,
	}
}