package handler

import (
	"context"
	"errors"

	"mail-service/internal/usecase"

	asosiyPB "github.com/khbdev/what-food-proto/proto/asosiy"
)

type FoodHandler struct {
	asosiyPB.UnimplementedFoodServiceServer
	usecase *usecase.FoodUsecase
}

func NewFoodHandler(u *usecase.FoodUsecase) *FoodHandler {
	return &FoodHandler{usecase: u}
}

// =========================
// FILTER FOOD
// =========================

func (h *FoodHandler) FilterFood(
	ctx context.Context,
	req *asosiyPB.FoodFilterRequest,
) (*asosiyPB.FoodListResponse, error) {

	if req.Country == "" {
		return nil, errors.New("country is required")
	}
	if req.MealTime == "" {
		return nil, errors.New("meal_time is required")
	}
	if req.KcalLimit <= 0 {
		return nil, errors.New("kcal_limit is invalid")
	}

	res, err := h.usecase.FilterFood(
		req.Country,
		req.MealTime,
		req.HasSalad,
		req.KcalLimit,
	)
	if err != nil {
		return nil, err
	}

	mailItems := make([]*asosiyPB.FoodItem, len(res.Items))
	for i, item := range res.Items {
		mailItems[i] = &asosiyPB.FoodItem{
			Id:           item.Id,
			Type:         item.Type,
			RestaurantId: item.RestaurantId,
			Name:         item.Name,
			Description:  item.Description,
			ImageUrl:     item.ImageUrl,
			VideoUrl:     item.VideoUrl,
			Country:      item.Country,
			MealTime:     item.MealTime,
			Kcal:         item.Kcal,
			Protein:      item.Protein,
			Fat:          item.Fat,
			Carbs:        item.Carbs,
		}
	}

	return &asosiyPB.FoodListResponse{Items: mailItems}, nil
}

// =========================
// FOOD DETAIL
// =========================

func (h *FoodHandler) GetFoodDetail(
	ctx context.Context,
	req *asosiyPB.FoodDetailRequest,
) (*asosiyPB.FoodDetailResponse, error) {

	if req.Id == 0 {
		return nil, errors.New("id is required")
	}
	if req.Type != "recipe" && req.Type != "salad" {
		return nil, errors.New("invalid type")
	}

	res, err := h.usecase.GetFoodDetailAndAnalyze(
		req.Id,
		req.Type,
		req.Portion,
		
	)
	if err != nil {
		return nil, err
	}

	var ingredients []*asosiyPB.Ingredient
	for _, ing := range res.Ingredients {
		ingredients = append(ingredients, &asosiyPB.Ingredient{
			Name:   ing.Name,
			Amount: ing.Amount,
		})
	}
return &asosiyPB.FoodDetailResponse{
    Id:                 res.Id,
    Type:               res.Type,
    RestaurantId:       res.RestaurantId,
    Name:               res.Name,
    Description:        res.Description,
    ImageUrl:           res.ImageUrl,
    VideoUrl:           res.VideoUrl,
    Country:            res.Country,
    MealTime:           res.MealTime,
    Kcal:               res.Kcal,
    Protein:            res.Protein,
    Fat:                res.Fat,
    Carbs:              res.Carbs,
    Portion:            res.Portion,
    TotalKcal:          res.TotalKcal,
    CookingTimeMinutes: res.CookingTimeMinutes,
    Ingredients:        ingredients,
    Steps:              res.Steps,
}, nil
}