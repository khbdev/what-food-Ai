package handler

import (
	"context"
	"errors"

	"mail-service/internal/usecase"

	foodpb  "github.com/khbdev/what-food-proto/proto/asosiy" 
)

// =========================
// gRPC HANDLER
// =========================

type FoodHandler struct {
	foodpb.UnimplementedFoodServiceServer

	usecase *usecase.FoodUsecase
}

func NewFoodHandler(u *usecase.FoodUsecase) *FoodHandler {
	return &FoodHandler{
		usecase: u,
	}
}

// =========================
// FILTER FOOD
// =========================

func (h *FoodHandler) FilterFood(
	ctx context.Context,
	req *foodpb.FoodFilterRequest,
) (*foodpb.FoodListResponse, error) {

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

	return &foodpb.FoodListResponse{
		Items: res.Items,
	}, nil
}

// =========================
// FOOD DETAIL
// =========================

func (h *FoodHandler) GetFoodDetail(
	ctx context.Context,
	req *foodpb.FoodDetailRequest,
) (*foodpb.FoodDetailResponse, error) {

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

	// ⚠️ AI response mapping (simple version)

	if req.Type == "recipe" {
		return &foodpb.FoodDetailResponse{
			Data: &foodpb.FoodDetailResponse_Recipe{
				Recipe: &foodpb.Recipe{
					Name:        "AI recipe result",
					Description: res.Steps[0],
				},
			},
		}, nil
	}

	return &foodpb.FoodDetailResponse{
		Data: &foodpb.FoodDetailResponse_Salad{
			Salad: &foodpb.Salad{
				Name:        "AI salad result",
				Description: res.Steps[0],
			},
		},
	}, nil
}