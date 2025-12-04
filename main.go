package main

import (
	"AuthInGo/app"
)

func main() {

	cfg := app.NewConfig(":3001")

	app := app.NewApplication(cfg)

	app.Run()
}
