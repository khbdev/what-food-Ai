package main

import (
	"log"
	"net"
	"os"

	"food-service/internal/config"
	"food-service/internal/handler"
	"food-service/internal/repository"
	"food-service/internal/usecase"
	// "food-service/pkg/env"

	foodpb "github.com/khbdev/what-food-proto/proto/food"
	"google.golang.org/grpc"
)

func main() {

	// // load env
	// env.LoadEnv()

	// db
	postgres, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	// redis
	redis, err := config.NewRedisClient()
	if err != nil {
		log.Fatal(err)
	}

	_ = redis

	// repositories
	recipeRepo := repository.NewRecipeRepository(postgres)
	saladRepo := repository.NewSaladRepository(postgres)
	filterRepo := repository.NewFoodFilterRepository(postgres)
	restaurantRepo := repository.NewRestaurantRepository(postgres)

	// usecases
	recipeUC := usecase.NewRecipeUsecase(recipeRepo)
	saladUC := usecase.NewSaladUsecase(saladRepo)
	filterUC := usecase.NewFoodFilterUsecase(filterRepo)
	restaurantUC := usecase.NewRestaurantUsecase(restaurantRepo)

	// handler
	hand := handler.NewFoodHandler(
		recipeUC,
		saladUC,
		restaurantUC,
		filterUC,
	)

	// grpc server
	grpcServer := grpc.NewServer()

	// register services
	foodpb.RegisterRecipeServiceServer(grpcServer, hand)
	foodpb.RegisterSaladServiceServer(grpcServer, hand)
	foodpb.RegisterRestaurantServiceServer(grpcServer, hand)
	foodpb.RegisterFoodFilterServiceServer(grpcServer, hand)

	// port from env
	port :=  os.Getenv("GRPC_PORT")

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("🚀 gRPC server running on port:", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}