package main

import (
	aimodel "ai-service/internal/ai-model"
	"ai-service/internal/handler"
	"ai-service/internal/usecase"
	"ai-service/pkg/env"
	pb "ai-service/proto"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	env.LoadEnv()

	groqAi := aimodel.NewGroqClient()
	aiUse := usecase.NewAIUsecase(groqAi)
	aiHand := handler.NewAIHandler(aiUse)

	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50055"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAiServiceServer(grpcServer, aiHand)

	log.Printf("AI Service running on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}