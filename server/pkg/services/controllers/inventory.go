package controllers

import (
	"context"
	"fmt"
	"time"

	col "github.com/ivankonevv/go_temi_rpc/internal/collections"
	pb "github.com/ivankonevv/go_temi_rpc/pkg/api/v1"
	"github.com/ivankonevv/go_temi_rpc/pkg/services/models"
	"github.com/ivankonevv/go_temi_rpc/platform/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateNewInventory(req *pb.NewInventoryRequest) (string, error) {
	var data []models.InventoryData

	for _, d := range req.Data {
		data = append(data, models.InventoryData{
			ID:     d.Id,
			Title:  d.Title,
			Serial: d.Serial,
			Color:  d.Color,
			Count:  d.Count,
		})
	}
	inventory := models.Inventory{
		Author:    req.Author,
		CreatedAt: time.Now().Local(),
		Data:      data,
	}
	valid, err := inventory.Validate()
	if err != nil || !valid {
		return "", status.Errorf(codes.InvalidArgument, fmt.Sprintf("Inventory validation failed: %v", err))
	}

	db := database.DB

	result, err := db.Collection(col.Inventories).InsertOne(context.Background(), &inventory)
	if err != nil {
		return "", status.Errorf(codes.Internal, fmt.Sprintf("Inventory insert failed: %v", err))
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}
