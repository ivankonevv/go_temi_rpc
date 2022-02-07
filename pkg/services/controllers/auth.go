package controllers

import (
	"context"
	"fmt"
	"temi_rpc/pkg/services/models"
	"temi_rpc/platform/database"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func FindUser(email string) (*models.User, error) {
	db := database.DB

	dbResult := models.User{}
	response := db.Collection("users").FindOne(context.Background(), bson.M{"email": email})
	if response.Err() != nil {
		logrus.Printf("Cannot find user with email `%s`", email)
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("User not found: %v", response.Err()))
	}
	if err := response.Decode(&dbResult); err != nil {
		logrus.Errorf("Cannot decode User result: %v", err)
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot decode result on FindUser: %v", err))
	}

	return &dbResult, nil
}
