package main

import (
	"log"
	"net"
	"os"

	"user-product-service/internal/cache"
	"user-product-service/internal/config"
	"user-product-service/internal/handler"
	repository "user-product-service/internal/repostory"
	"user-product-service/internal/usecase"
	"user-product-service/pkg/loadenv"

	incrideatspb "github.com/khbdev/what-food-proto/proto/incrideats"
	categorypb "github.com/khbdev/what-food-proto/proto/products"
	"google.golang.org/grpc"
)

func main() {
	loadenv.LoadEnv()

	postgress, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	redis, err := config.NewRedisClient()
	if err != nil {
		log.Fatal(err)
	}

	repoCategory := repository.NewCategoryRepository(postgress)
	repoIncrideat := repository.NewIngredientRepository(postgress)

	cacheCategory := cache.NewCategoryCache(redis)

	srvCategory := usecase.NewCategoryUsecase(repoCategory, cacheCategory)
	srvProduct := usecase.NewIngredientUsecase(repoIncrideat)

	handCategory := handler.NewCategoryHandler(srvCategory)
	handProduct := handler.NewProductsHandler(srvProduct)

	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	incrideatspb.RegisterIngredientServiceServer(grpcServer, handProduct)
	incrideatspb.RegisterCategoryServiceServer(grpcServer, handCategory)

	log.Printf("gRPC server started on port %s", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}