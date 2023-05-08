package services

import (
	"context"
	"grpc_go/configs"
	pb "grpc_go/proto"
)

var db = configs.NewDBHandler()

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (service UserServiceServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	resp, err := db.GetUser(req.Id)

	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{Id: resp.Id.String(), Name: resp.Name, Location: resp.Location, Title: resp.Title}, nil
}

func (service UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	newUser := configs.User{Name: req.Name, Location: req.Location, Title: req.Title}
	_, err := db.CreateUser(newUser)

	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{Data: "User created successfully!"}, nil
}

func (service UserServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	newUser := configs.User{Name: req.Name, Location: req.Location, Title: req.Title}
	_, err := db.UpdateUser(req.Id, newUser)

	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{Data: "User updated successfully!"}, nil
}

func (service UserServiceServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	_, err := db.DeleteUser(req.Id)

	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{Data: "User details deleted successfully!"}, nil
}

func (service UserServiceServer) GetAllUsers(context.Context, *pb.Empty) (*pb.GetAllUsersResponse, error) {
	resp, err := db.GetAllUsers()
	var users []*pb.UserResponse

	if err != nil {
		return nil, err
	}

	for _, v := range resp {
		var singleUser = &pb.UserResponse{
			Id:       v.Id.String(),
			Name:     v.Name,
			Location: v.Location,
			Title:    v.Title,
		}
		users = append(users, singleUser)
	}

	return &pb.GetAllUsersResponse{Users: users}, nil
}
