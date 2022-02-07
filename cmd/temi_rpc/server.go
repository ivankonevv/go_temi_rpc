package main

import (
	"net"
	"os"
	"temi_rpc/internal/roles"
	"temi_rpc/pkg/api/v1"
	handlers2 "temi_rpc/pkg/services/handlers"
	"temi_rpc/platform/database"
	"temi_rpc/tools"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Errorf("Cannot find .env file: %v", err)
	}
	database.Connect()

	jwtManager := tools.NewJWTManager(os.Getenv("JWT_SECRET_KEY"), time.Hour*10)

	interceptor := tools.NewAuthInterceptor(jwtManager, roles.AccessibleRoles())
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)
	listener, err := net.Listen("tcp", os.Getenv("SERVER_URL"))
	if err != nil {
		logrus.Errorf("Cannot create net listener: %v", err)
	}
	v1.RegisterMetalDoorsApiServer(grpcServer, &handlers2.DoorsApiServer{})
	logrus.Info("MetalDoors service connected...")
	v1.RegisterWoodDoorsApiServer(grpcServer, &handlers2.WoodDoorsApiService{})
	logrus.Info("WoodDoors service connected...")
	v1.RegisterColorsApiServer(grpcServer, &handlers2.WoodColorsApiService{})
	logrus.Info("Colors service connected...")
	v1.RegisterUserServiceServer(grpcServer, &handlers2.UserServer{})
	authServer := handlers2.NewAuthServer(jwtManager)

	v1.RegisterAuthServiceServer(grpcServer, authServer)

	v1.RegisterInventoryServiceServer(grpcServer, &handlers2.InventoryServer{})
	v1.RegisterCallbackServiceServer(grpcServer, &handlers2.CallbackServer{})

	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		logrus.Errorf("Cannot start rpc server: %v", err)
	}
}
