package models

import (
	pb "temi_rpc/pkg/api/v1"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SingleMetalDoorResponse struct {
	ID             primitive.ObjectID                   `bson:"_id" json:"id"`
	Title          string                               `bson:"title" json:"title"`
	Price          map[string]float32                   `bson:"price" json:"price"`
	Specifications *pb.Specifications                   `bson:"specifications" json:"specifications"`
	Variants       []*pb.SingleMetalDoorResponseVariant `bson:"variants" json:"variants"`
	Tags           []string                             `bson:"tags" json:"tags"`
}

type Variant struct {
	InColor   string   `bson:"in_color" json:"in_color"`
	OutColor  string   `bson:"out_color" json:"out_color"`
	InImages  []string `bson:"in_images" json:"in_images"`
	OutImages []string `bson:"out_images" json:"out_images"`
	Size      []Size   `bson:"size" json:"size"`
	Hex       string   `bson:"hex" json:"hex"`
}

type Size struct {
	Width int32      `bson:"width" json:"width"`
	Side  string     `bson:"side" json:"side"`
	Data  []SizeData `bson:"data" json:"data"`
}

type SizeData struct {
	Serial         string `bson:"serial" json:"serial"`
	Dealer         string `bson:"dealer" json:"dealer"`
	State          string `bson:"state" json:"state"`
	Prepay         string `bson:"prepay" json:"prepay"`
	OrderDate      string `bson:"order_date" json:"order_date"`
	ReadyDate      string `bson:"ready_date" json:"ready_date"`
	ReserveEndDate string `bson:"reserve_end_date" json:"reserve_end_date"`
}

type WriteMetalDoor struct {
	UUID           string             `bson:"uuid" json:"uuid"`
	Title          string             `bson:"title" json:"title"`
	ForCatalog     bool               `bson:"for_catalog" json:"for_catalog"`
	Tags           []string           `bson:"tags" json:"tags"`
	Price          map[string]float32 `bson:"price" json:"price"`
	Variants       []*pb.Variant      `bson:"variants" json:"variants"`
	Specifications *pb.Specifications `bson:"specifications" json:"specifications"`
}

// repeated WoodDoorVariant variants = 8;
// repeated Reserve reserves = 9;

type GetMetalDoor struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	UUID           string             `bson:"uuid" json:"uuid"`
	Title          string             `bson:"title" json:"title"`
	ForCatalog     bool               `bson:"for_catalog" json:"for_catalog"`
	Tags           []string           `bson:"tags" json:"tags"`
	Price          map[string]float32 `bson:"price" json:"price"`
	Variants       []*pb.Variant      `bson:"variants" json:"variants"`
	Specifications *pb.Specifications `bson:"specifications" json:"specifications"`
}

type Specifications struct {
	Height     float32 `bson:"height" json:"height"`         // Высота
	Metal      float32 `bson:"metal" json:"metal"`           // Толщина металла
	Leaf       float32 `bson:"leaf" json:"leaf"`             // Толщина полотна
	UpperLock  string  `bson:"upper_lock" json:"upper_lock"` // Верхний замок
	LowerLock  string  `bson:"lower_lock" json:"lower_lock"` // Нижний замок
	Valve      string  `bson:"valve" json:"valve"`           // Задвижка
	Armor      string  `bson:"armor" json:"armor"`           // Броне накладка
	Cylinder   string  `bson:"cylinder" json:"cylinder"`     // Цилиндр
	Peephole   string  `bson:"peephole" json:"peephole"`     // Глазок
	Insulation string  `bson:"insulation" json:"insulation"` // Утеплитель
	Sealer     string  `bson:"sealer" json:"sealer"`         // Уплотнитель
	Handle     string  `bson:"handle" json:"handle"`         // Ручка
}

type Post struct {
	ID        string             `bson:"_id"`
	Title     string             `bson:"title"`
	Price     map[string]float32 `bson:"price"`
	Tags      []string           `bson:"tags"`
	InColors  []string           `bson:"in_colors"`
	InImages  []string           `bson:"in_images"`
	OutImages []string           `bson:"out_images"`
}

// string id = 1;
// string title = 2;
// float price = 3;
// repeated string tags = 4;
// repeated InColors in_colors = 5;
// repeated string in_images = 6;
// repeated string out_images = 7;
