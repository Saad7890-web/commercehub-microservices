package grpc

import (
	"fmt"
	"log"
	"net"

	"auth-service/auth-service/pkg/pb"
	"auth-service/internal/usecase"

	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	port       string
	authUC     usecase.AuthUsecase
}

func NewServer(port string, authUC usecase.AuthUsecase) *Server {
	return &Server{
		grpcServer: grpc.NewServer(),
		port:       port,
		authUC:     authUC,
	}
}

func (s *Server) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	handler := NewAuthHandler(s.authUC)
	pb.RegisterAuthServiceServer(s.grpcServer, handler)

	log.Println("gRPC server listening on port", s.port)
	if err := s.grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
