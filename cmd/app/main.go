package main

import (
	"debts-service/config"
	"debts-service/internal/app"
)

func main() {
	cfg := config.NewConfig()

	app.Run(&cfg)
}
