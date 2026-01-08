package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"commercehub/services/auth/internal/config"
	"commercehub/services/auth/internal/infrastructure/postgres"
	authgrpc "commercehub/services/auth/internal/transport/grpc"
)

func main() {
	cfg := config.Load()

	log.Printf("starting %s", cfg.ServiceName)

	dbPool := postgres.NewPool(cfg.DBURL)
	defer dbPool.Close()

	grpcServer := authgrpc.New(cfg.GRPCPort)

	go grpcServer.Start()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown
	log.Println("shutdown signal received")

	grpcServer.Stop()
	log.Println("auth service stopped gracefully")
}
