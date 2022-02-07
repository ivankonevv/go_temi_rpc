package controllers

import (
	"context"
	"fmt"
	pb "temi_rpc/pkg/api/v1"
	"temi_rpc/pkg/services/models"
	"temi_rpc/platform/database"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetUserData(email string) (*models.GetUser, error) {
	db := database.DB

	result := db.Collection("users").FindOne(context.Background(), bson.M{"email": email})
	if result.Err() != nil {
		return nil, status.Errorf(codes.NotFound, "requested user not found")
	}
	user := models.GetUser{}
	err := result.Decode(&user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("cannot decore db_helpers response: %v", err))
	}

	return &user, nil
}

func GetUsers() ([]*pb.GetUserResponse, error) {
	db := database.DB

	dbResult := &models.GetUser{}
	cursor, err := db.Collection("users").Find(
		context.Background(),
		bson.M{},
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			logrus.Errorf("GetUsers: cannot close cursor: %v", err)
			return
		}
	}(cursor, context.Background())

	if cursor.Err() != nil {
		logrus.Infof("GetUsers: users not found: %v", err)
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Users not found: %v", err))
	}
	var result []*pb.GetUserResponse
	for cursor.Next(context.Background()) {
		if err := cursor.Decode(&dbResult); err != nil {
			logrus.Errorf("Cannot decode result: %v", err)
			return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Cannot decode result: %v", err))
		}

		result = append(result, &pb.GetUserResponse{
			Id:        dbResult.ID.Hex(),
			FirstName: dbResult.FirstName,
			LastName:  dbResult.LastName,
			Email:     dbResult.Email,
			Role:      dbResult.Role,
		})
	}

	return result, nil
}

func CreateNewUser(req *pb.NewUserRequest) (string, error) {
	user := models.User{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      req.Role,
	}
	valid, err := user.Validate(req.Password)
	if err != nil || !valid {
		return "", status.Errorf(codes.InvalidArgument, fmt.Sprintf("user is not valid: %v", err))
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return "", status.Errorf(codes.Internal, fmt.Sprintf("cannot generate hash password: %v", err))
	}
	user.HashedPassword = string(password)

	db := database.DB
	result, err := db.Collection("users").InsertOne(context.Background(), &user)
	if err != nil {
		return "", status.Errorf(codes.Internal, fmt.Sprintf("cannot insert user: %v", err))
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}
