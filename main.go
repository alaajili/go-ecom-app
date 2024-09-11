package main

import (
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api"
	"log"
)

func main() {

	cfg, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	api.StartServer(cfg)

}