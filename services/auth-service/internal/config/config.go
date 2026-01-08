package config

import (
	"log"
	"os"
)

type Config struct {
	ServiceName string
	GRPCPort    string
	DBURL       string
}

func Load() *Config {
	cfg := &Config{
		ServiceName: getEnv("SERVICE_NAME", "auth-service"),
		GRPCPort:    getEnv("GRPC_PORT", "50051"),
		DBURL:       getEnv("DATABASE_URL", ""),
	}

	if cfg.DBURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
