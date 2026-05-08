package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"analytics-service/internal/config"
	"analytics-service/internal/handler"
	repository "analytics-service/internal/repostory"
	"analytics-service/internal/usecase"
	loadenv "analytics-service/pkg/load_env"

	pb "nutrition/proto"
)

func main() {

	// =====================
	// LOAD ENV
	// =====================
	loadenv.LoadEnv()

	// =====================
	// DB
	// =====================
	db, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	// =====================
	// RABBITMQ
	// =====================
	rabbit := config.NewRabbit()

	// =====================
	// REPOSITORY
	// =====================
	repo := repository.NewMealRepository(db)

	// =====================
	// USECASE
	// =====================
	use := usecase.NewMealUsecase(repo)

	// =====================
	// HANDLERS
	// =====================
	consumerHandler := handler.NewHandlerConsumer(rabbit.Channel, use)
	nutritionHandler := handler.NewNutritionHandler(use)

	// background consumer
	go consumerHandler.Start()

	// =====================
	// gRPC SERVER
	// =====================a
	lis, err := net.Listen("tcp", config.GRPCPort())
	if err != nil {
		log.Fatal("failed to listen:", err)
	}

	grpcServer := grpc.NewServer()

	// register service
	pb.RegisterNutritionServiceServer(grpcServer, nutritionHandler)

	log.Println("🚀 server running on", config.GRPCPort())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}