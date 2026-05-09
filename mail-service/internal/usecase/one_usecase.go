package usecase

import (




	foodpb "github.com/khbdev/what-food-proto/proto/food"
	aipb "github.com/khbdev/what-food-proto/proto/ai"
)

// =========================
// USECASE
// =========================

type FoodUsecase struct {
	foodClient *client.FoodClient
	aiClient   *cli.AiClient
}

// =========================
// INIT
// =========================

func NewFoodUsecase(f *client.FoodClient, ai *client.AiClient) *FoodUsecase {
	return &FoodUsecase{
		foodClient: f,
		aiClient:   ai,
	}
}