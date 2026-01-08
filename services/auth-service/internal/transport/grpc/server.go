package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	listener   net.Listener
}

func New(port string) *Server {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	return &Server{
		grpcServer: s,
		listener:   lis,
	}
}

func (s *Server) Start() {
	log.Println("gRPC server started")
	if err := s.grpcServer.Serve(s.listener); err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}
}

func (s *Server) Stop() {
	log.Println("stopping gRPC server")
	s.grpcServer.GracefulStop()
}
