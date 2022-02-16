package handlers

import (
	"context"

	"github.com/ivankonevv/go_temi_rpc/pkg/api/v1"
	"github.com/ivankonevv/go_temi_rpc/pkg/services/controllers"
	"github.com/ivankonevv/go_temi_rpc/tools"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	v1.UnimplementedAuthServiceServer
	jwtManager *tools.JWTManager
}

func NewAuthServer(jwtManager *tools.JWTManager) *AuthServer {
	return &AuthServer{v1.UnimplementedAuthServiceServer{}, jwtManager}
}

func (server *AuthServer) Login(_ context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	user, err := controllers.FindUser(req.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := server.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create token: %v", err)
	}

	res := &v1.LoginResponse{AccessToken: token}
	return res, nil
}
