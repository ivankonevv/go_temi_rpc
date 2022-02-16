package models

import (
	pb "github.com/ivankonevv/go_temi_rpc/pkg/api/v1"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetColor struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	UUID         string             `bson:"uuid" json:"uuid"`
	Title        string             `bson:"title" json:"title"`
	HEX          string             `bson:"hex" json:"hex"`
	RAL          string             `bson:"ral" json:"ral"`
	Image        string             `bson:"image" json:"image"`
	Thumbnail    string             `bson:"thumbnail" json:"thumbnail"`
	Manufacturer string             `bson:"manufacturer" json:"manufacturer"`
	Collection   *pb.Collection     `bson:"collection" json:"collection"`
}

type WriteColor struct {
	UUID         string         `bson:"uuid" json:"uuid"`
	Title        string         `bson:"title" json:"title"`
	HEX          string         `bson:"hex" json:"hex"`
	RAL          string         `bson:"ral" json:"ral"`
	Image        string         `bson:"image" json:"image"`
	Thumbnail    string         `bson:"thumbnail" json:"thumbnail"`
	Manufacturer string         `bson:"manufacturer" json:"manufacturer"`
	Collection   *pb.Collection `bson:"collection" json:"collection"`
}

type GetMwColors struct {
	ID           primitive.ObjectID                              `bson:"_id" json:"id"`
	Manufacturer string                                          `bson:"manufacturer"`
	Collections  []*pb.ManufacturersWColorsResponse_MCollections `bson:"collections" json:"collections"`
}
