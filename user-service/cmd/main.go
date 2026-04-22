package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"user-service/internal/cache"
	"user-service/internal/config"
	"user-service/internal/handler"
	"user-service/internal/usecase"

	repository "user-service/internal/repostory"
	loadenv "user-service/pkg/loadEnv"

	userpb "github.com/khbdev/what-food-proto/proto/userr"
)

func main() {
	// load env
	loadenv.LoadEnv()

	// config
	db, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	redisClient, err := config.NewRedisClient()
	if err != nil {
		log.Fatal(err)
	}

	// layers
	userCache := cache.NewUserCache(redisClient)
	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUsecase(userRepo, userCache)
	userHandler := handler.NewUserHandler(userUC)

	// grpc server
	grpcServer := grpc.NewServer()

	// register service
	userpb.RegisterUserServiceServer(grpcServer, userHandler)

	// port from env (default 50050)
	port := os.Getenv("GRPC_PORT")

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("failed to listen:", err)
	}

	log.Println("🚀 gRPC server running on port:", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve:", err)
	}
}