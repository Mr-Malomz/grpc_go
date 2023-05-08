package services

import (
	"context"
	"grpc_go/configs"
	pb "grpc_go/proto"
)

var db = configs.NewDBHandler()

type UserServiceServer struct{}

func GetUser(context.Context, *pb.UserRequest) (*pb.UserResponse, error) {
	
}
