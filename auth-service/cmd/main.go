package main

import (
	"log"

	"auth-service/internal/config"
	grpcDelivery "auth-service/internal/delivery/grpc"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	cfg := config.Load()

	server := grpcDelivery.NewServer(cfg.ServicePort)
	server.Start()
}
