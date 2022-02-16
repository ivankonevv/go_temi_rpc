package models

import (
	"fmt"
	"time"

	pb "github.com/ivankonevv/go_temi_rpc/pkg/api/v1"
)

type Inventory struct {
	Author    string          `json:"author" bson:"author"`
	Mentor    string          `json:"mentor" bson:"mentor"`
	CreatedAt time.Time       `json:"created_at" bson:"created_at"`
	CheckedAt time.Time       `json:"checked_at" bson:"checked_at"`
	Valid     bool            `json:"valid" bson:"valid"`
	Data      []InventoryData `json:"data" bson:"data"`
}

type InventoryData struct {
	ID     string                           `json:"id" bson:"door_id"`
	Title  string                           `json:"title" bson:"title"`
	Serial string                           `json:"serial" bson:"serial"`
	Color  string                           `json:"color" bson:"color"`
	Count  []*pb.InventoryModel_SizeCounter `json:"count" bson:"count"`
}

type InventoryDataCount struct {
	Width float32 `json:"width" bson:"width"`
	Side  string  `json:"side" bson:"side"`
	Count int32   `json:"count" bson:"count"`
	Valid bool    `json:"is_valid" bson:"is_valid"`
}

func (i *Inventory) Validate() (bool, error) {
	if len(i.Data) < 1 {
		return false, fmt.Errorf("inventory data must have at least one object")
	}
	if i.Author == "" {
		return false, fmt.Errorf("inventory author is not defined")
	}
	return true, nil
}
