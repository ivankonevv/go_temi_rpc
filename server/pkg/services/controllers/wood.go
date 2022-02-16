package controllers

import (
	"context"
	"fmt"

	col "github.com/ivankonevv/go_temi_rpc/internal/collections"
	pb "github.com/ivankonevv/go_temi_rpc/pkg/api/v1"
	pipelines "github.com/ivankonevv/go_temi_rpc/pkg/db_helpers"
	"github.com/ivankonevv/go_temi_rpc/pkg/services/models"
	"github.com/ivankonevv/go_temi_rpc/platform/database"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetWoodDoors() ([]*pb.WoodDoorsResponse, error) {
	db := database.DB

	dbResult := &models.GetWoodDoors{}

	cursor, err := db.Collection(col.WoodDoors).Aggregate(context.Background(), pipelines.ProductWColorsPipeline)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			logrus.Errorf("GetDoors: cannot close cursor: %v", err)
			return
		}
	}(cursor, context.Background())
	if cursor.Err() != nil {
		logrus.Infof("GetWoodDoors: posts not found: %v", err)
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Posts not found: %v", err))
	}
	var result []*pb.WoodDoorsResponse
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&dbResult)
		if err != nil {
			return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decoed data: %v", err))
		}
		var colors []*pb.WSColor
		var images []string
		for _, v := range dbResult.Variants {
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
		result = append(result, &pb.WoodDoorsResponse{
			Data: &pb.WoodDoor{
				Id:     dbResult.ID.Hex(),
				Title:  dbResult.Title,
				Price:  dbResult.Price,
				Tags:   dbResult.Tags,
				Colors: colors,
				Images: images,
			},
		})
	}
	return result, nil
}
