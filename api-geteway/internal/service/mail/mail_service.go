package service

import (
	"context"
	"errors"
	"strings"

	
	asosiypb "github.com/khbdev/what-food-proto/proto/asosiy"
)

type FoodService struct {
	foodClient *.FoodClient
}

// =========================
// INIT
// =========================

func NewFoodService(c *client.FoodClient) *FoodService {
	return &FoodService{
		foodClient: c,
	}
}

// =========================
// FILTER FOOD
// =========================

func (s *FoodService) FilterFood(ctx context.Context, req *asosiypb.FoodFilterRequest) (*asosiypb.FoodListResponse, error) {

	if strings.TrimSpace(req.Country) == "" {
		return nil, errors.New("country is required")
	}

	if strings.TrimSpace(req.MealTime) == "" {
		return nil, errors.New("meal_time is required")
	}

	if req.KcalLimit < 0 {
		return nil, errors.New("kcal_limit cannot be negative")
	}

	return s.foodClient.FilterFood(req)
}

// =========================
// GET FOOD DETAIL
// =========================

func (s *FoodService) GetFoodDetail(ctx context.Context, req *asosiypb.FoodDetailRequest) (*asosiypb.FoodDetailResponse, error) {

	if req.Id == 0 {
		return nil, errors.New("id is required")
	}

	if strings.TrimSpace(req.Type) == "" {
		return nil, errors.New("type is required")
	}

	if req.Portion <= 0 {
		return nil, errors.New("portion must be greater than 0")
	}

	return s.foodClient.GetFoodDetail(req)
}