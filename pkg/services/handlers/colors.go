package handlers

import (
	"context"
	"fmt"
	"temi_rpc/pkg/api/v1"
	"temi_rpc/pkg/db_helpers"
	"temi_rpc/pkg/services/models"
	database2 "temi_rpc/platform/database"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WoodColorsApiService struct {
	v1.UnimplementedColorsApiServer
}

type GetCreator interface {
	GetColor() error
}

// TODO: refactor this into separate files

func (s *WoodColorsApiService) GetColors(req *v1.ColorsRequest, stream v1.ColorsApi_GetColorsServer) error {
	db := database2.DB
	result := models.GetColor{}
	cur, err := db.Collection("wood-colors").Find(context.Background(), bson.M{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			return
		}
	}(cur, context.Background())
	for cur.Next(context.Background()) {
		err = cur.Decode(&result)
		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Cannot decode result: %v", err))
		}

		err = stream.Send(&v1.ColorsResponse{
			Data: &v1.WSColor{
				Id:           result.ID.Hex(),
				Uuid:         result.UUID,
				Image:        result.Image,
				Thumbnail:    result.Thumbnail,
				Ral:          result.RAL,
				Title:        result.Title,
				Hex:          result.HEX,
				Manufacturer: result.Manufacturer,
				Collection:   result.Collection,
			},
		})
		if err != nil {
			return status.Errorf(codes.NotFound, fmt.Sprintf("Strem sending error: %v", err))
		}
	}

	return nil
}

func (s *WoodColorsApiService) CreateColor(ctx context.Context, req *v1.NewColorRequest) (*v1.NewColorResponse, error) {
	db := database2.DB
	color, err := db.Collection("wood-colors").InsertOne(ctx, models.WriteColor{
		Title:        req.Title,
		UUID:         uuid.New().String(),
		HEX:          req.Hex,
		RAL:          req.Ral,
		Image:        req.Image,
		Thumbnail:    req.Thumbnail,
		Manufacturer: req.Manufacturer,
		Collection:   req.Collection,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot insert data into WColors: %v", err))
	}

	id := color.InsertedID.(primitive.ObjectID).Hex()
	return &v1.NewColorResponse{Id: id}, nil
}

func (s *WoodColorsApiService) GetMwColors(req *v1.ManufacturersWColorsRequest, stream v1.ColorsApi_GetMwColorsServer) error {
	db := database2.DB
	result := models.GetMwColors{}
	cur, err := db.Collection("wood-colors-manufacturers").Aggregate(context.Background(), db_helpers.MWCCollectionPipeline)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			return
		}
	}(cur, context.Background())
	for cur.Next(context.Background()) {
		err = cur.Decode(&result)
		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Cannot decode result: %v", err))
		}
		fmt.Println(result)
		err = stream.Send(&v1.ManufacturersWColorsResponse{
			Id:           result.ID.Hex(),
			Manufacturer: result.Manufacturer,
			Collections:  result.Collections,
		})
		if err != nil {
			return status.Errorf(codes.NotFound, fmt.Sprintf("Strem sending error: %v", err))
		}
	}

	return nil
}
