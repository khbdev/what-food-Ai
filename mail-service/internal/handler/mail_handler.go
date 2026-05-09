package handler

import (
	"context"
	"errors"

	"mail-service/internal/usecase"

	foodpb "github.com/khbdev/what-food-proto/proto/food"
)

// =========================
// gRPC HANDLER
// =========================

type FoodGRPCHandler struct {
	foodpb.UnimplementedFoodServiceServer

	usecase *usecase.FoodUsecase
}

func NewFoodGRPCHandler(u *usecase.FoodUsecase) *FoodGRPCHandler {
	return &FoodGRPCHandler{
		usecase: u,
	}
}

// =========================
// FILTER FOOD
// =========================

func (h *FoodGRPCHandler) FilterFood(
	ctx context.Context,
	req *foodpb.FoodFilterRequest,
) (*foodpb.FoodListResponse, error) {

	res, err := h.usecase.FilterFood(
		req.Country,
		req.MealTime,
		req.IncludeSalads,
		req.,
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

func (h *FoodGRPCHandler) GetFoodDetail(
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

	// ⚠️ AI response -> hozircha simple mapping
	// real loyihada bu yerda full mapping qilinadi

	if req.Type == "recipe" {
		return &foodpb.FoodDetailResponse{
			Data: &foodpb.FoodDetailResponse_Recipe{
				Recipe: &foodpb.Recipe{
					Name:        "AI processed recipe",
					Description: res.Steps[0],
				},
			},
		}, nil
	}

	return &foodpb.FoodDetailResponse{
		Data: &foodpb.FoodDetailResponse_Salad{
			Salad: &foodpb.Salad{
				Name:        "AI processed salad",
				Description: res.Steps[0],
			},
		},
	}, nil
}