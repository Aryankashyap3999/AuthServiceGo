package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func GetString(key string, fallback string) string {
	load()

	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return value
}

func GetInt(key string, fallback int) int {

	load()

	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Error converting env variable to int:", err)
		return fallback
	}

	return intValue
}

func GetBool(key string, fallback bool) bool {
	load()

	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		fmt.Println("Error converting env variable to bool:", err)
		return fallback
	}

	return boolValue
}

