package usecase

import (
	"errors"

	"auth-service/internal/domain"
	"auth-service/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthUsecase {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) Register(email, password string) error {
	existingUser, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user := &domain.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.userRepo.Create(user)
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	); err != nil {
		return "", errors.New("invalid credentials")
	}

	// JWT will be added in next step
	return "dummy-token", nil
}
