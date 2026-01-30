package grpc

import (
	"context"

	"auth-service/auth-service/pkg/pb"
	"auth-service/internal/usecase"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

func (h *AuthHandler) Register(
	ctx context.Context,
	req *pb.RegisterRequest,
) (*pb.RegisterResponse, error) {

	err := h.authUsecase.Register(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Message: "user registered successfully",
	}, nil
}

func (h *AuthHandler) Login(
	ctx context.Context,
	req *pb.LoginRequest,
) (*pb.LoginResponse, error) {

	token, err := h.authUsecase.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		AccessToken: token,
	}, nil
}
