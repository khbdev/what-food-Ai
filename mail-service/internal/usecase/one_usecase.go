package usecase

import (
	"context"
	"errors"
	"time"

	"mail-service/internal/client"
	"mail-service/internal/config"
	"mail-service/internal/models"
	rabbitmqproducer "mail-service/pkg/rabbitmq_producer"

	aipb "github.com/khbdev/what-food-proto/proto/ai"
	foodpb "github.com/khbdev/what-food-proto/proto/food"
)

type FoodUsecase struct {
	foodClient *client.FoodClient
	aiClient   *client.AiClient
	rabbit     *config.Rabbit
	mealCache *
}

func NewFoodUsecase(
	f *client.FoodClient,
	ai *client.AiClient,
	r *config.Rabbit,
) *FoodUsecase {
	return &FoodUsecase{
		foodClient: f,
		aiClient:   ai,
		rabbit:     r,
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
	userID uint,
) (*DetailResult, error) {
 
	if id == 0 {
		return nil, errors.New("id is required")
	}
 
	if foodType != "recipe" && foodType != "salad" {
		return nil, errors.New("invalid type")
	}
 
	ctx := context.Background()
 
	// =========================
	// READ-THROUGH: Redis dan qidiramiz
	// Cache hit bo'lsa — food fetch ham, AI ham o'tkazib yuboriladi
	// =========================
 
	cached, err := u.mealCache.Get(ctx, foodType, id)
	if err != nil {
		// Redis ishlamayapti — loglab davom etamiz
		// log.Printf("cache get error (id=%d): %v", id, err)
	}
 
	if cached != nil {
		// ✅ Cache HIT — hech qanday external request ketmaydi
		result := &DetailResult{
			// Food ma'lumotlari
			Id:           cached.Id,
			Type:         cached.Type,
			RestaurantId: cached.RestaurantId,
			Name:         cached.Name,
			Description:  cached.Description,
			ImageUrl:     cached.ImageUrl,
			VideoUrl:     cached.VideoUrl,
			Country:      cached.Country,
			MealTime:     cached.MealTime,
			Kcal:         cached.Kcal,
			Protein:      cached.Protein,
			Fat:          cached.Fat,
			Carbs:        cached.Carbs,
			// AI ma'lumotlari
			Portion:            cached.Portion,
			TotalKcal:          cached.TotalKcal,
			CookingTimeMinutes: cached.CookingTimeMinutes,
			Ingredients:        cached.Ingredients,
			Steps:              cached.Steps,
		}
 
		go u.publishMealEvent(userID, result)
 
		return result, nil
	}
 
	// =========================
	// Cache MISS — Food fetch + AI request
	// =========================
 
	result := &DetailResult{
		Id:   id,
		Type: foodType,
	}
 
	// FOOD FETCH
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
 
	// AI REQUEST
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
 
	// =========================
	// WRITE-THROUGH: to'liq natijani Redis ga saqlaymiz
	// =========================
 
	toCache := &cache.CachedMealAnalysis{
		// Food ma'lumotlari
		Id:           result.Id,
		Type:         result.Type,
		RestaurantId: result.RestaurantId,
		Name:         result.Name,
		Description:  result.Description,
		ImageUrl:     result.ImageUrl,
		VideoUrl:     result.VideoUrl,
		Country:      result.Country,
		MealTime:     result.MealTime,
		Kcal:         result.Kcal,
		Protein:      result.Protein,
		Fat:          result.Fat,
		Carbs:        result.Carbs,
		// AI ma'lumotlari
		Portion:            result.Portion,
		TotalKcal:          result.TotalKcal,
		CookingTimeMinutes: result.CookingTimeMinutes,
		Ingredients:        result.Ingredients,
		Steps:              result.Steps,
	}
 
	if err := u.mealCache.Set(ctx, foodType, id, toCache); err != nil {
		// Redis yozish xatosi — kritik emas, faqat loglaymiz
		// log.Printf("cache set error (id=%d): %v", id, err)
	}
 
	// =========================
	// RABBIT MQ — async event
	// =========================
 
	go u.publishMealEvent(userID, result)
 
	return result, nil
}
 
// publishMealEvent - RabbitMQ ga async yuborish, alohida ajratildi (DRY)
func (u *FoodUsecase) publishMealEvent(userID uint, result *DetailResult) {
	mealEvent := &models.Meal{
		UserID:   userID,
		Name:     result.Name,
		Country:  result.Country,
		MealTime: time.Now(),
		Kcal:     float64(result.TotalKcal),
		Protein:  result.Protein,
		Fat:      result.Fat,
		Carbs:    result.Carbs,
	}
 
	if err := rabbitmqproducer.PublishMeal(u.rabbit.Channel, mealEvent); err != nil {
		// log.Println("rabbit publish error:", err)
	}
}