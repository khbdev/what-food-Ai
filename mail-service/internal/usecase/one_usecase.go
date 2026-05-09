package usecase

import (
	"errors"

	"mail-service/internal/client"

	aipb "github.com/khbdev/what-food-proto/proto/ai"
	foodpb "github.com/khbdev/what-food-proto/proto/food"
)

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
// FILTER RESULT
// =========================

type FilterResult struct {
	Items []*foodpb.FoodItem
}

// =========================
// DETAIL RESULT
// =========================

type DetailResult struct {
	// food ma'lumotlari
	Id           int64
	Type         string
	RestaurantId int64
	Name         string
	Description  string
	ImageUrl     string
	VideoUrl     string
	Country      string
	MealTime     string
	Kcal         int32
	Protein      float64
	Fat          float64
	Carbs        float64

	// AI ma'lumotlari
	Portion            int32
	TotalKcal          float32
	CookingTimeMinutes int32
	Ingredients        []*aipb.Ingredient
	Steps              []string
}

// =========================
// 1. FILTER FOOD
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
		Country:       country,
		MealTime:      mealTime,
		IncludeSalads: hasSalad,
		MaxKcal:       maxKcal,
	}

	items, err := u.foodClient.FilterFood(req)
	if err != nil {
		return nil, err
	}

	return &FilterResult{Items: items}, nil
}

// =========================
// 2. DETAIL + AI ANALYZE
// =========================

func (u *FoodUsecase) GetFoodDetailAndAnalyze(
	id int64,
	foodType string,
	portion int32,
) (*DetailResult, error) {

	if id == 0 {
		return nil, errors.New("id is required")
	}
	if foodType != "recipe" && foodType != "salad" {
		return nil, errors.New("invalid type")
	}

	result := &DetailResult{
		Id:   id,
		Type: foodType,
	}

	// =========================
	// GET FOOD BY TYPE
	// =========================

	if foodType == "recipe" {
		r, err := u.foodClient.GetRecipeByID(id)
		if err != nil {
			return nil, err
		}
		result.RestaurantId = r.RestaurantId
		result.Name = r.Name
		result.Description = r.Description
		result.ImageUrl = r.ImageUrl
		result.VideoUrl = r.VideoUrl
		result.Country = r.Country
		result.MealTime = r.MealTime
		result.Kcal = r.Kcal
		result.Protein = r.Protein
		result.Fat = r.Fat
		result.Carbs = r.Carbs
	}

	if foodType == "salad" {
		s, err := u.foodClient.GetSaladByID(id)
		if err != nil {
			return nil, err
		}
		result.RestaurantId = s.RestaurantId
		result.Name = s.Name
		result.Description = s.Description
		result.ImageUrl = s.ImageUrl
		result.VideoUrl = s.VideoUrl
		result.Country = s.Country
		result.MealTime = s.MealTime
		result.Kcal = s.Kcal
		result.Protein = s.Protein
		result.Fat = s.Fat
		result.Carbs = s.Carbs
	}

	// =========================
	// AI REQUEST
	// =========================

	aiReq := &aipb.MealRequest{
		Name:        result.Name,
		Description: result.Description,
		Country:     result.Country,
		MealTime:    result.MealTime,
		Kcal:        float32(result.Kcal),
		Protein:     float32(result.Protein),
		Fat:         float32(result.Fat),
		Carbs:       float32(result.Carbs),
		Portion:     portion,
	}

	aiRes, err := u.aiClient.AnalyzeMeal(aiReq)
	if err != nil {
		return nil, err
	}

	result.Portion = aiRes.Portion
	result.TotalKcal = aiRes.TotalKcal
	result.CookingTimeMinutes = aiRes.CookingTimeMinutes
	result.Ingredients = aiRes.Ingredients
	result.Steps = aiRes.Steps

	return result, nil
}