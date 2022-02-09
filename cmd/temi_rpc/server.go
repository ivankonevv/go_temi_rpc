package main

import (
	"net"
	"os"
	"temi_rpc/internal/roles"
	"temi_rpc/pkg/api/v1"
	"temi_rpc/pkg/services/handlers"
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

	// backendCert, _ := ioutil.ReadFile(os.Getenv("SSL_CERT"))
	// backendKey, _ := ioutil.ReadFile(os.Getenv("SSL_KEY"))
	// cert, err := tls.X509KeyPair(backendCert, backendKey)
	// if err != nil {
	// 	logrus.Errorf("Cannot create SSL certificate")
	// }
	// creds := credentials.NewServerTLSFromCert(&cert)

	jwtManager := tools.NewJWTManager(os.Getenv("JWT_SECRET_KEY"), time.Hour*10)

	interceptor := tools.NewAuthInterceptor(jwtManager, roles.AccessibleRoles())
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
		// grpc.Creds(creds),
	)
	listener, err := net.Listen("tcp", os.Getenv("SERVER_URL"))
	if err != nil {
		logrus.Errorf("Cannot create net listener: %v", err)
	}
	v1.RegisterMetalDoorsApiServer(grpcServer, &handlers.DoorsApiServer{})
	logrus.Info("MetalDoors service connected...")
	v1.RegisterWoodDoorsApiServer(grpcServer, &handlers.WoodDoorsApiService{})
	logrus.Info("WoodDoors service connected...")
	v1.RegisterColorsApiServer(grpcServer, &handlers.WoodColorsApiService{})
	logrus.Info("Colors service connected...")
	v1.RegisterUserServiceServer(grpcServer, &handlers.UserServer{})
	authServer := handlers.NewAuthServer(jwtManager)

	v1.RegisterAuthServiceServer(grpcServer, authServer)

	v1.RegisterInventoryServiceServer(grpcServer, &handlers.InventoryServer{})
	v1.RegisterCallbackServiceServer(grpcServer, &handlers.CallbackServer{})

	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		logrus.Errorf("Cannot start rpc server: %v", err)
	}
}
