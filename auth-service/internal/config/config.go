package config

import (
	"log"
	"os"
)

type Config struct {
	ServiceName string
	ServicePort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecret string
}

func Load() *Config {
	cfg := &Config{
		ServiceName: getEnv("SERVICE_NAME", "auth-service"),
		ServicePort: getEnv("SERVICE_PORT", "50051"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "auth_db"),

		JWTSecret: getEnv("JWT_SECRET", "secret"),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("env %s not set, using default", key)
		return fallback
	}
	return value
}
