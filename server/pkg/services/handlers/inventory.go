package handlers

import (
	"context"

	pb "github.com/ivankonevv/go_temi_rpc/pkg/api/v1"
	"github.com/ivankonevv/go_temi_rpc/pkg/services/controllers"

	"github.com/sirupsen/logrus"
)

type InventoryServer struct {
	pb.UnimplementedInventoryServiceServer
}

func (server *InventoryServer) NewInventory(_ context.Context, req *pb.NewInventoryRequest) (*pb.NewInventoryResponse, error) {

	id, err := controllers.CreateNewInventory(req)
	if err != nil {
		logrus.Errorf("%v", err)
		return nil, err
	}
	return &pb.NewInventoryResponse{Id: id}, nil
}
