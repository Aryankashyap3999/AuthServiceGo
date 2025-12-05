package db

import (
	// "database/sql"
	"fmt"
)

type UsersRepository interface { // faciliates dependancy injection interface for repository
	Create() error
}

type UserRepositoryImp struct {
	// db *sql.DB
}

func (u *UserRepositoryImp) Create()  error {
	fmt.Println("Creating user in repository layer")	
	return nil
}

func NewUserRepository() UsersRepository {
	return &UserRepositoryImp{}
}