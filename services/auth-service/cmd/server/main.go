package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	grpcPort = ":50051"
)

func main(){
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", grpcPort, err)
	}

	grpcServer :=grpc.NewServer()
	log.Printf("Auth Service gRPC server running on %s", grpcPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
