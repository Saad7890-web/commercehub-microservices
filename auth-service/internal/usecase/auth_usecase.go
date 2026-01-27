package usecase

type AuthUsecase interface {
	Register(email, password string) error
	Login(email, password string) (string, error)
}
