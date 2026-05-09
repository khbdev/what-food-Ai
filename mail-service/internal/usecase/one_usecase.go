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


type FilterResult struct {
	Portion int32
	Items   []*foodpb.FoodItem
}

func (u *FoodUsecase) FilterFood(req *foodpb.FoodFilterRequest) (*FilterResult, error) {

	// =========================
	// VALIDATION
	// =========================

	if req.Country == "" {
		return nil, er.New("country is required")
	}

	if req.MealTime == "" {
		return nil, errors.New("meal_time is required")
	}

	if req.MaxKcal <= 0 {
		return nil, errors.New("max_kcal is invalid")
	}

	// =========================
	// CALL CLIENT
	// =========================

	items, err := u.foodClient.FilterFood(req)
	if err != nil {
		return nil, err
	}

	// =========================
	// RETURN CONTEXT (portion saqlanadi)
	// =========================

	return &FilterResult{
		Portion: req.Portion,
		Items:   items,
	}, nil
}