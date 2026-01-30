package postgres

import (
	"database/sql"
	"errors"

	"auth-service/internal/domain"
	"auth-service/internal/repository"

	"github.com/google/uuid"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
	user.ID = uuid.NewString()

	query := `
		INSERT INTO users (id, email, password)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(query, user.ID, user.Email, user.Password)
	return err
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	query := `
		SELECT id, email, password
		FROM users
		WHERE email = $1
	`

	row := r.db.QueryRow(query, email)

	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
