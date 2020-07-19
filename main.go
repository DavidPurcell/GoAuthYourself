package main

import (
	"github.com/DavidPurcell/GoAuthYourself/app"
	"github.com/DavidPurcell/GoAuthYourself/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
