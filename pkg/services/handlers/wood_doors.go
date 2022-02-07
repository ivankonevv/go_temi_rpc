package handlers

import (
	"context"
	"fmt"
	pb "temi_rpc/pkg/api/v1"
	"temi_rpc/pkg/db_helpers"
	"temi_rpc/pkg/services/models"
	database2 "temi_rpc/platform/database"

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
	return &pb.CreateWoodDoorResponse{Status: "success", Id: &id}, nil
}

func (s *WoodDoorsApiService) GetWoodDoors(req *pb.WoodDoorsRequest, stream pb.WoodDoorsApi_GetWoodDoorsServer) error {
	db := database2.DB
	result := &models.GetWoodDoors{}

	cur, err := db.Collection("wood-doors").Aggregate(context.Background(), db_helpers.ProductWColorsPipeline)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		err := cur.Decode(result)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decoed data: %v", err))
		}
		var colors []*pb.WSColor
		var images []string
		for _, v := range result.Variants {
			colors = append(colors, &pb.WSColor{
				Id:           v.Color.Id,
				Uuid:         v.Color.Uuid,
				Image:        v.Color.Image,
				Thumbnail:    v.Color.Thumbnail,
				Ral:          v.Color.Ral,
				Title:        v.Color.Title,
				Hex:          v.Color.Hex,
				Manufacturer: v.Color.Manufacturer,
			})
			for _, i := range v.Images {
				images = append(images, i)
			}
		}
		stream.Send(&pb.WoodDoorsResponse{
			Data: &pb.WoodDoor{
				Id:     result.ID.Hex(),
				Title:  result.Title,
				Price:  result.Price,
				Tags:   result.Tags,
				Colors: colors,
				Images: images,
			},
		})
	}
	if err := cur.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	return nil
}
