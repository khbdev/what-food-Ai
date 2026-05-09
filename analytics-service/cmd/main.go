package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"analytics-service/internal/config"
	"analytics-service/internal/handler"
	repository "analytics-service/internal/repostory"
	"analytics-service/internal/usecase"
	loadenv "analytics-service/pkg/load_env"

pb  "github.com/khbdev/what-food-proto/proto/statik"
)

func main() {

	// =====================
	// LOAD ENV
	// =====================
	loadenv.LoadEnv()

	// =====================
	// PORT (DIRECT ENV)
	// =====================
	port := os.Getenv("PORT")
	if port == "" {
		port = ":50056"
	}

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
	// =====================
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen:", err)
	}

	grpcServer := grpc.NewServer()

	// register service
	pb.RegisterNutritionServiceServer(grpcServer, nutritionHandler)

	log.Println("🚀 gRPC server running on", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}