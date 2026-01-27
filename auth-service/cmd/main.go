package main

import (
	"fmt"
	"log"

	"auth-service/internal/config"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	cfg := config.Load()

	fmt.Println("Starting service:", cfg.ServiceName)
	fmt.Println("Listening on port:", cfg.ServicePort)
}
