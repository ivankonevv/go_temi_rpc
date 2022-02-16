package models

import (
	"github.com/ivankonevv/go_temi_rpc/pkg/api/v1"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WriteWoodDoor struct {
	UUID           string                     `bson:"uuid" json:"uuid"`
	Title          string                     `bson:"title" json:"title"`
	ForCatalog     bool                       `bson:"for_catalog" json:"for_catalog"`
	Tags           []string                   `bson:"tags" json:"tags"`
	Price          map[string]float32         `bson:"price" json:"price"`
	Manufacturer   string                     `bson:"manufacturer" json:"manufacturer"`
	Collection     string                     `bson:"collection" json:"collection"`
	Specifications *v1.WoodDoorSpecifications `bson:"specifications" json:"specifications"`
	Variants       []*v1.WoodDoorVariant      `bson:"variants" json:"variants"`
	Reserves       []*v1.Reserve              `bson:"reserves" json:"reserves"`
}

type GetWoodDoors struct {
	ID             primitive.ObjectID         `bson:"_id" json:"id"`
	UUID           string                     `bson:"uuid" json:"uuid"`
	Title          string                     `bson:"title" json:"title"`
	ForCatalog     bool                       `bson:"for_catalog" json:"for_catalog"`
	Tags           []string                   `bson:"tags" json:"tags"`
	Price          float32                    `bson:"price" json:"price"`
	Variants       []GetWoodDoorsVariant      `bson:"variants" json:"variants"`
	Specifications *v1.WoodDoorSpecifications `bson:"specifications" json:"specifications"`
}

type GetWoodDoorsVariant struct {
	Color  *v1.WSColor `bson:"color" json:"color"`
	Images []string    `bson:"images" json:"images"`
}

// message WSColor {
// string id = 1;
// string uuid = 2;
// string color = 3;
// string ral = 4;
// string title = 5;
// string hex = 6;
// string manufacturer = 7;
// }
