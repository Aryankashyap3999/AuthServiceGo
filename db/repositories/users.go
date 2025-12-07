package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UsersRepository interface { // faciliates dependancy injection interface for repository
	GetById() (*models.User, error)
	Create(username string, email string, hashedpassword string) (*models.User, error)
	GetAll() ([]*models.User, error)
	GetByEmail(email string) (*models.User, error)
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

func (u *UserRepositoryImp) Create(username string, email string, hashedpassword string) (*models.User, error) {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

	result, err := u.db.Exec(query, username, email, hashedpassword)

	user := &models.User{}



	if err != nil {
		fmt.Println("Error inserting user:", err)
		return nil, err
	}

	rowsAffected, rowErr := result.RowsAffected()

	if rowErr != nil {
		fmt.Println("Error fetching rows affected:", rowErr)
		return nil, rowErr
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected, user not created")
		return nil, nil
	}

	user.Id, _ = result.LastInsertId()
	user.Username = username
	user.Email = email


	fmt.Println("User successfully created with ID:", user.Id)

	return user,nil
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
	query := "SELECT id, username, email, created_at, updated_at FROM users"
	rows, err := u.db.Query(query)
	if err != nil {
		fmt.Println("Error fetching users:", err)
		return nil, err
	}

	defer rows.Close() // Ensure rows are closed after processing

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			fmt.Println("Error scanning user:", err)
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error with rows:", err)
		return nil, err
	}

	return users, nil

}

func (u *UserRepositoryImp) GetByEmail(email string) (*models.User, error) {
	query := "SELECT id, email, password FROM users WHERE email = ?"
	row := u.db.QueryRow(query, email)
	users := &models.User{}

	err := row.Scan(&users.Id, &users.Email, &users.Password)
	
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given email")
			return nil, err
		} else {
			fmt.Println("Error scanning user by email:", err)
			return nil, err
		}
	}

	return users, nil
}

func (u *UserRepositoryImp) DeleteById(id int64) error {
	query := "DELETE FROM users WHERE id = ?"

	result, err := u.db.Exec(query, id)

	if err != nil {
		fmt.Println("Error deleting user:", err)
		return err
	}

	rowsAffected, rowErr := result.RowsAffected()

	if rowErr != nil {
		fmt.Println("Error fetching rows affected:", rowErr)
		return rowErr
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected, user not deleted")
		return nil
	}

	fmt.Println("User successfully deleted", rowsAffected)

	return nil
}