package controllers

import (
	"context"
	"fmt"
	pb "temi_rpc/pkg/api/v1"
	"temi_rpc/pkg/db_helpers"
	"temi_rpc/pkg/services/models"
	database2 "temi_rpc/platform/database"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetSingleDoor func gets a single product from db.
// Return codes.NotFound if product with such id is not exists, codes.Internal if response from db can't be decoded.
func GetSingleDoor(id primitive.ObjectID) (*pb.SingleMetalDoorResponse, error) {
	db := database2.DB

	dbResult := &models.SingleMetalDoorResponse{}
	res := db.Collection("metal-doors").FindOne(context.Background(), bson.M{"_id": id})
	if res.Err() != nil {
		logrus.Printf("Cannot find post with id `%s`", id.Hex())
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Post not found: %v", res.Err()))
	}
	if err := res.Decode(&dbResult); err != nil {
		logrus.Errorf("Cannot decode SingleMetalDoor result: %v", err)
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot decode result on GetSingleDoor: %v", err))
	}
	result := &pb.SingleMetalDoorResponse{
		Id:             dbResult.ID.Hex(),
		Title:          dbResult.Title,
		Price:          dbResult.Price["roz"],
		Specifications: dbResult.Specifications,
		Variants:       dbResult.Variants,
		Tags:           dbResult.Tags,
	}

	return result, nil
}

// GetDoors func gets a list of product from db.
func GetDoors() ([]*pb.PostsResponse, error) {
	db := database2.DB

	dbResult := &models.GetMetalDoor{}
	cursor, err := db.Collection("metal-doors").Find(
		context.Background(),
		bson.M{},
		options.Find().SetProjection(db_helpers.GetDoorsProjection),
	)
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
		logrus.Infof("GetDoors: posts not found: %v", err)
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Posts not found: %v", err))
	}
	var result []*pb.PostsResponse
	for cursor.Next(context.Background()) {
		if err := cursor.Decode(&dbResult); err != nil {
			logrus.Errorf("Cannot decode result: %v", err)
			return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Cannot decode result: %v", err))
		}
		var inColors []*pb.InColors
		var inImages []string
		var outImages []string
		for _, v := range dbResult.Variants {
			inColors = append(inColors, &pb.InColors{
				Color: v.InColor,
				Hex:   v.Hex,
			})
			inImages = append(inImages, v.InImages[0])
			outImages = append(outImages, v.OutImages[0])
		}
		result = append(result, &pb.PostsResponse{
			Data: &pb.Post{
				Id:        dbResult.ID.Hex(),
				Title:     dbResult.Title,
				Price:     dbResult.Price["roz"],
				Tags:      dbResult.Tags,
				InColors:  inColors,
				InImages:  inImages,
				OutImages: outImages,
			},
		})
	}

	return result, nil
}

// CreateDoor - creates a new post with req data. Returns created object id and error.
func CreateDoor(req *pb.CreatePostRequest) (primitive.ObjectID, error) {
	db := database2.DB

	post, err := db.Collection("metal-doors").InsertOne(context.Background(), models.WriteMetalDoor{
		UUID:           uuid.New().String(),
		Title:          req.Title,
		ForCatalog:     req.ForCatalog,
		Tags:           req.Tags,
		Price:          req.Price,
		Variants:       req.Variants,
		Specifications: req.Specifications,
	})
	if err != nil {
		logrus.Errorf("Cannot insert post: %v", err)
		return primitive.NilObjectID, status.Errorf(codes.Internal, fmt.Sprintf("Cannot insert post: %v", err))
	}

	id := post.InsertedID.(primitive.ObjectID)

	return id, nil
}
