package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func New(cfg Config) *Postgres {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to open database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping database:", err)
	}

	log.Println("PostgreSQL connected")

	return &Postgres{DB: db}
}
