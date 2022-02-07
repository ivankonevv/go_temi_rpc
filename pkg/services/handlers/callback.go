package handlers

import (
	"context"
	"fmt"
	"temi_rpc/pkg/api/v1"
	"temi_rpc/pkg/services/controllers"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CallbackServer struct {
	v1.UnimplementedCallbackServiceServer
}

func (server *CallbackServer) Callback(_ context.Context, req *v1.CallbackRequest) (*v1.CallbackResponse, error) {
	body := fmt.Sprintf("%s\n%s\n", req.GetName(), req.GetPhone())
	if err := controllers.SendMail(body); err != nil {
		return &v1.CallbackResponse{Status: v1.CallbackResponse_ERROR}, status.Errorf(codes.Internal, fmt.Sprintf("error sending callback request: %v", err))
	}

	return &v1.CallbackResponse{Status: v1.CallbackResponse_SUCCESS}, nil
}
