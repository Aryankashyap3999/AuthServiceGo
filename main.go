package main

import (
	"AuthInGo/app"
	dbConfig "AuthInGo/config/db"
)

func main() {

	cfg := app.NewConfig()

	app := app.NewApplication(cfg)
	dbConfig.SetupDB()
	app.Run()
}
