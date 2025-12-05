package config

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"

	env "AuthInGo/config/env"
)

func SetupDB() (*sql.DB, error){
	cfg := mysql.NewConfig()
	cfg.User = env.GetString("DB_USER", "root")
	cfg.Passwd = env.GetString("DB_PASSWORD", "password")
	cfg.Net = env.GetString("DB_NET", "tcp")
	cfg.Addr = env.GetString("DB_ADDR", "localhost:3306")
	cfg.DBName = env.GetString("DBName", "auth_dev")

	fmt.Println("Database configuration:", cfg.FormatDSN())

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil, err
	}

	fmt.Println("Connecting to database...")
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Error pinging database:", pingErr)
		return nil, pingErr
	}	

	fmt.Println("Successfully connected to database!", cfg.DBName)

	return db, nil
}



