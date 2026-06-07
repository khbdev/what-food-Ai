package main

import (
	"fmt"
	"log"
	"net"
	repository "notifaction-service/internal/repostory"
	"os"

	"notifaction-service/internal/config"
	"notifaction-service/internal/handler"

	"notifaction-service/internal/usecase"
	"notifaction-service/pkg/loadenv"
	redis_pub "notifaction-service/pkg/redis-pub"

	pb "github.com/khbdev/what-food-proto/proto/notifaction-crud"
	"google.golang.org/grpc"
)

func main() {

	// 1. env load
	loadenv.LoadEnv()

	// 2. port
	port := getEnv("GRPC_PORT", "50059")

	// 3. infra
	db := config.NewPostgresConnection()
	redisClient := config.InitRedis()
	rabbit := config.NewRabbit()

	// 4. repos + pub/sub
	pubRedis := redis_pub.NewNotificationPublisher(redisClient)
	repo := repository.NewNotificationRepository(db)

	// 5. usecase
	notificationUC := usecase.NewNotificationUsecase(repo, pubRedis)
	smsUC := usecase.NewSMSUsecase()

	// 6. handler
	notificationHandler := handler.NewNotificationHandler(notificationUC)
	consumerHandler := handler.NewHandlerConsumer(rabbit.Channel, smsUC)

	// 7. gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// register gRPC service
	pb.RegisterNotificationServiceServer(grpcServer, notificationHandler)

	// start consumer (async)
	go consumerHandler.Start()

	fmt.Println("gRPC server running on port:", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
