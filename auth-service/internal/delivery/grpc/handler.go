package grpc

import (
	"context"

	"auth-service/auth-service/pkg/pb"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Register(
	ctx context.Context,
	req *pb.RegisterRequest,
) (*pb.RegisterResponse, error) {
	return &pb.RegisterResponse{
		Message: "register endpoint not implemented",
	}, nil
}

func (h *AuthHandler) Login(
	ctx context.Context,
	req *pb.LoginRequest,
) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{
		AccessToken: "",
	}, nil
}
