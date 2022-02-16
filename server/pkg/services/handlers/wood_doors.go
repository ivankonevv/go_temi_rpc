package handlers

import (
	"context"
	"fmt"

	pb "github.com/ivankonevv/go_temi_rpc/pkg/api/v1"
	"github.com/ivankonevv/go_temi_rpc/pkg/services/controllers"
	"github.com/ivankonevv/go_temi_rpc/pkg/services/models"
	database2 "github.com/ivankonevv/go_temi_rpc/platform/database"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WoodDoorsApiService struct {
	pb.UnimplementedWoodDoorsApiServer
}

// TODO: refactor this into separate files

func (s *WoodDoorsApiService) CreateWoodDoor(ctx context.Context, req *pb.CreateWoodDoorRequest) (*pb.CreateWoodDoorResponse, error) {
	db := database2.DB
	post, err := db.Collection("wood-doors").InsertOne(ctx, models.WriteWoodDoor{
		UUID:           uuid.New().String(),
		Title:          req.Title,
		ForCatalog:     req.ForCatalog,
		Tags:           req.Tags,
		Price:          req.Price,
		Manufacturer:   req.Manufacturer,
		Collection:     req.Collection,
		Specifications: req.Specifications,
		Variants:       req.Variants,
		Reserves:       req.Reserves,
	})
	if err != nil {
		logrus.Warnf("Cannot insert post: %v", err)
		return &pb.CreateWoodDoorResponse{Status: "error"}, err
	}

	id := post.InsertedID.(primitive.ObjectID).Hex()
	return &pb.CreateWoodDoorResponse{Status: "success", Id: id}, nil
}

func (s *WoodDoorsApiService) GetWoodDoors(_ *pb.WoodDoorsRequest, stream pb.WoodDoorsApi_GetWoodDoorsServer) error {
	result, err := controllers.GetWoodDoors()
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
