package main

import (
	"sppg-backend/config"
	"sppg-backend/internal/app"
)

func main() {
	config.NewConfig()
	app.Run()
}