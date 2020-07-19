package main

import (
	"github.com/davidpurcell/GoAuthYourself/app"
	"github.com/davidpurcell/GoAuthYourself/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
