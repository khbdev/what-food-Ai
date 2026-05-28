package main

import (
	"log"
	"net"
	"os"

	"mail-service/internal/client"
	"mail-service/internal/config"
	"mail-service/internal/handler"
	"mail-service/internal/usecase"
	// "mail-service/pkg/cache"
	"mail-service/pkg/loadenv"

	asosiyPB "github.com/khbdev/what-food-proto/proto/asosiy"
	feedbackPB "github.com/khbdev/what-food-proto/proto/feedback"
	"google.golang.org/grpc"
)

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("❌ %s is empty", key)
	}
	return v
}

func main() {
loadenv.LoadEnv()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	rabbitMq := config.NewRabbit()
	redis, err := config.NewRedisClient()
	if err != nil {
		log.Fatal(err)
	}


	// cacheRedis := cache.NewAIMealCache(redis)

	// _ = cacheRedis

	aiURL := mustEnv("AI_URL")
	foodURL := mustEnv("FOOD_URL")

	aiClient, err := client.NewAiClient(aiURL)
	if err != nil {
		log.Fatalf("❌ AI client error: %v", err)
	}
	defer aiClient.Close()

	foodClient, err := client.NewFoodClient(foodURL)
	if err != nil {
		log.Fatalf("❌ Food client error: %v", err)
	}
	defer foodClient.Close()

	uc := usecase.NewFoodUsecase(foodClient, aiClient, rabbitMq, redis)
	ucFeedBack := usecase.NewNutritionUsecase(aiClient)
	h := handler.NewFoodHandler(uc)
	hFeedback := handler.NewNutritionHandler(ucFeedBack)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("❌ Listen error: %v", err)
	}

	grpcServer := grpc.NewServer()
	asosiyPB.RegisterFoodServiceServer(grpcServer, h)
	feedbackPB.RegisterMailServiceServer(grpcServer, hFeedback)


	log.Printf("🚀 mail-service running on :%s", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Serve error: %v", err)
	}
}