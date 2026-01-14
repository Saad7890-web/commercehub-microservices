package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"user-service/internal/handler"
	"user-service/internal/repository"
	"user-service/internal/service"
)

func main() {
	db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/user_service?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	r := mux.NewRouter()
	r.HandleFunc("/users", h.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", h.GetUser).Methods("GET")

	log.Println("User service running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
