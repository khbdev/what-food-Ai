package usecase

import (
	"mail-service/internal/client"

	aipb "github.com/khbdev/what-food-proto/proto/ai"
	foodpb "github.com/khbdev/what-food-proto/proto/food"
)

// =========================
// USECASE
// =========================

type FoodUsecase struct {
	foodClient *client.FoodClient
	aiClient   *client.AiClient
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

