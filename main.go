package main

import (
	"log"
	"net"

	pb "grpc_go/proto"
	"grpc_go/services"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "[::1]:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	service := &services.UserServiceServer{}

	pb.RegisterUserServiceServer(grpcServer, service)
	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalf("Error strating server: %v", err)
	}
}
