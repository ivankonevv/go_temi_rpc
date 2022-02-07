package handlers

import (
	"context"
	"fmt"
	"temi_rpc/pkg/api/v1"
	"temi_rpc/pkg/services/controllers"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DoorsApiServer struct {
	v1.UnimplementedMetalDoorsApiServer
}

func (s *DoorsApiServer) CreatePost(_ context.Context, req *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {
	id, err := controllers.CreateDoor(req)
	if err != nil {
		return nil, err
	}

	sId := id.Hex()

	return &v1.CreatePostResponse{Id: &sId}, nil
}

func (s *DoorsApiServer) GetPosts(_ *v1.PostsRequest, stream v1.MetalDoorsApi_GetPostsServer) error {
	result, err := controllers.GetDoors()
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

func (s *DoorsApiServer) GetSingleMetalDoor(_ context.Context, req *v1.SingleMetalDoorRequest) (*v1.SingleMetalDoorResponse, error) {
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		logrus.Errorf("GetSingleMetalDoor: cannot convert id to Object: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot convert id to Object: %v", err))
	}
	result, err := controllers.GetSingleDoor(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
