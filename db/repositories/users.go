package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UsersRepository interface { // faciliates dependancy injection interface for repository
	GetById() (*models.User, error)
	Create(username string, email string, hashedpassword string) (error)
	GetAll() ([]*models.User, error)
	DeleteById(id int64) error
}

type UserRepositoryImp struct {
	db *sql.DB
}

func NewUserRepository(_db *sql.DB) UsersRepository {
	return &UserRepositoryImp{
		db: _db,
	}
}

func (u *UserRepositoryImp) Create(username string, email string, hashedpassword string) (error) {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

	result, err := u.db.Exec(query, username, email, hashedpassword)

	if err != nil {
		fmt.Println("Error inserting user:", err)
		return err
	}

	rowsAffected, rowErr := result.RowsAffected()

	if rowErr != nil {
		fmt.Println("Error fetching rows affected:", rowErr)
		return rowErr
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected, user not created")
		return nil
	}

	fmt.Println("User successfully created", rowsAffected)

	return nil
}

func (u *UserRepositoryImp) GetById()  (*models.User, error) {
	fmt.Println("Fetching user in repository layer")	

	query := "SELECT id, username, email, created_at, updated_at FROM users WHERE id = ?"
	row := u.db.QueryRow(query, 1)

	user := &models.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given ID")
			return nil, err
		} else {
			fmt.Println("Error scanning user:", err)		
			return nil, err
		}
	}

	fmt.Printf("User fetched: %+v\n", user)

	return user, nil
}

func (u *UserRepositoryImp) GetAll() ([]*models.User, error) {
	return nil, nil
}

func (u *UserRepositoryImp) DeleteById(id int64) error {
	return nil
}