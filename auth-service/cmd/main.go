package main

import (
	"log"
	"net"
	"os"

	"auth-service/internal/client"
	"auth-service/internal/config"
	"auth-service/internal/handler"
	"auth-service/internal/usecase"

	authpb "github.com/khbdev/what-food-proto/proto/auth"

	loadenv "auth-service/pkg/LoadEnv"
	rabbitMq "auth-service/pkg/rabbitMq"
	Redis "auth-service/pkg/redis"

	"google.golang.org/grpc"
)
func main() {
	loadenv.Load()

	

	// rabbit
	rabbitmq := config.NewRabbit()

	// redis
	redisClient, err := config.NewRedisClient()
	if err != nil {
		log.Fatal("redis error:", err)
	}

	// publisher
	producer := rabbitMq.NewPublisher(rabbitmq)

	

	// user service client (gRPC)
	userClient, err := client.NewUserClient(os.Getenv("USER_SERVICE"))
	if err != nil {
		log.Fatal("user client error:", err)
	}

	// redis service wrapper
	rd := Redis.NewService(redisClient)
	phoneCache := Redis.NewPhoneCache(redisClient)

	// usecase
	usc := usecase.NewAuthUsecase(userClient, rd, producer, phoneCache)

	// handler
	hand := handler.NewAuthHandler(usc)

	// gRPC server
	server := grpc.NewServer()

	// register service
	authpb.RegisterAuthServiceServer(server, hand)

	// port from env
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("listen error:", err)
	}

	log.Println("🚀 Auth gRPC server running on port:", port)

	// serve
	if err := server.Serve(lis); err != nil {
		log.Fatal("server error:", err)
	}
}