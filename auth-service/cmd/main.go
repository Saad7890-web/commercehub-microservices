package main

import (
	"log"

	"auth-service/internal/config"
	grpcDelivery "auth-service/internal/delivery/grpc"
	"auth-service/internal/infrastructure/postgres"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	cfg := config.Load()

	pg := postgres.New(postgres.Config{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
	})

	_= pg

	server := grpcDelivery.NewServer(cfg.ServicePort)
	server.Start()
}
