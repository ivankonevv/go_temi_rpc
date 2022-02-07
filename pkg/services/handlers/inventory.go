package handlers

import (
	"context"
	"temi_rpc/pkg/api/v1"
	"temi_rpc/pkg/services/controllers"

	"github.com/sirupsen/logrus"
)

type InventoryServer struct {
	v1.UnimplementedInventoryServiceServer
}

func (server *InventoryServer) NewInventory(ctx context.Context, req *v1.NewInventoryRequest) (*v1.NewInventoryResponse, error) {

	id, err := controllers.CreateNewInventory(req)
	if err != nil {
		logrus.Errorf("%v", err)
		return nil, err
	}
	return &v1.NewInventoryResponse{Id: id}, nil
}
