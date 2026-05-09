package usecase

import (
	"errors"

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

func NewFoodUsecase(f *client.FoodClient, ai *client.AiClient) *FoodUsecase {
	return &FoodUsecase{
		foodClient: f,
		aiClient:   ai,
	}
}

// =========================
// FILTER RESULT WRAPPER
// =========================

type FilterResult struct {
	Items []*foodpb.FoodItem
}

// =========================
// 1. FILTER FOOD (FOOD SERVICE)
// =========================

func (u *FoodUsecase) FilterFood(
	country string,
	mealTime string,
	hasSalad bool,
	maxKcal int32,
) (*FilterResult, error) {

	if country == "" {
		return nil, errors.New("country is required")
	}

	if mealTime == "" {
		return nil, errors.New("meal_time is required")
	}

	if maxKcal <= 0 {
		return nil, errors.New("max_kcal is invalid")
	}

	req := &foodpb.FoodFilterRequest{
		Country:    country,
		MealTime:   mealTime,
		IncludeSalads:   hasSalad,
		MaxKcal:    maxKcal,
	}

	items, err := u.foodClient.FilterFood(req)
	if err != nil {
		return nil, err
	}

	return &FilterResult{
		Items: items,
	}, nil
}

// =========================
// 2. DETAIL + AI ANALYZE
// =========================

func (u *FoodUsecase) GetFoodDetailAndAnalyze(
	id int64,
	foodType string,
	portion int32,
) (*aipb.MealResponse, error) {

	if id == 0 {
		return nil, errors.New("id is required")
	}

	if foodType != "recipe" && foodType != "salad" {
		return nil, errors.New("invalid type")
	}

	var (
		name, desc, country, mealTime string
		kcal, protein, fat, carbs      float64
	)

	// =========================
	// GET FOOD BY TYPE
	// =========================

	if foodType == "recipe" {

		r, err := u.foodClient.GetRecipeByID(id)
		if err != nil {
			return nil, err
		}

		name = r.Name
		desc = r.Description
		country = r.Country
		mealTime = r.MealTime
		kcal = float64(r.Kcal)
		protein = r.Protein
		fat = r.Fat
		carbs = r.Carbs
	}

	if foodType == "salad" {

		s, err := u.foodClient.GetSaladByID(id)
		if err != nil {
			return nil, err
		}

		name = s.Name
		desc = s.Description
		country = s.Country
		mealTime = s.MealTime
		kcal = float64(s.Kcal)
		protein = s.Protein
		fat = s.Fat
		carbs = s.Carbs
	}

	// =========================
	// AI REQUEST
	// =========================

	aiReq := &aipb.MealRequest{
		Name:        name,
		Description: desc,
		Country:     country,
		MealTime:    mealTime,
		Kcal:        float32(kcal),
		Protein:     float32(protein),
		Fat:         float32(fat),
		Carbs:       float32(carbs),
		Portion:     portion,
	}

	return u.aiClient.AnalyzeMeal(aiReq)
}