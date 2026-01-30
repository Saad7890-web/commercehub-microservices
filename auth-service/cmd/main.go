package main

import (
	"log"

	"auth-service/internal/config"
	grpcDelivery "auth-service/internal/delivery/grpc"
	"auth-service/internal/infrastructure/postgres"
	"auth-service/internal/usecase"

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

	userRepo := postgres.NewUserRepository(pg.DB)
	authUC := usecase.NewAuthService(userRepo)

	server := grpcDelivery.NewServer(cfg.ServicePort, authUC)
	server.Start()
}
