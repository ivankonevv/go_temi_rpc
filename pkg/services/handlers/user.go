package handlers

import (
	"context"
	"fmt"
	pb "temi_rpc/pkg/api/v1"
	"temi_rpc/pkg/services/controllers"
	"temi_rpc/tools"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (server *UserServer) GetUser(ctx context.Context, _ *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is missing")
	}

	accessToken := values[0]
	claims, err := tools.ExtractTokenClaims(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "cannot extract token claims")
	}
	result, err := controllers.GetUserData(claims.Email)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		Id:        result.ID.Hex(),
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Email:     result.Email,
		Role:      result.Role,
	}, nil
}

func (server *UserServer) GetAllUsers(_ *pb.GetAllUsersRequest, stream pb.UserService_GetAllUsersServer) error {
	result, err := controllers.GetUsers()
	if err != nil {
		return err
	}
	for _, v := range result {
		if err := stream.Send(v); err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Internal server error: %v", err))
		}
	}
	return nil
}

func (server *UserServer) CreateUser(_ context.Context, req *pb.NewUserRequest) (*pb.NewUserResponse, error) {
	if len(req.Email) < 4 && len(req.Password) < 8 {
		return nil, status.Errorf(codes.Aborted, fmt.Sprintf("email must be at least 4 characters and password must be at least 8 characters"))
	}

	id, err := controllers.CreateNewUser(req)
	if err != nil {
		return nil, err
	}

	return &pb.NewUserResponse{Id: id}, nil
}
