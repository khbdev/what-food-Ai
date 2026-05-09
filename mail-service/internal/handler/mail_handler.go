package handler

import (
	"context"
	"errors"

	"mail-service/internal/usecase"

	asosiyPB "github.com/khbdev/what-food-proto/proto/asosiy"
)

// =========================
// gRPC HANDLER
// =========================

type FoodHandler struct {
	asosiyPB.UnimplementedFoodServiceServer

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

	return &asosiyPB.FoodListResponse{
		Items: mailItems,
	}, nil
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

	if req.Type == "recipe" {
		return &asosiyPB.FoodDetailResponse{
			Data: &asosiyPB.FoodDetailResponse_Recipe{
				Recipe: &asosiyPB.Recipe{
					Name:        "AI recipe result",
					Description: res.Steps[0],
				},
			},
		}, nil
	}

	return &asosiyPB.FoodDetailResponse{
		Data: &asosiyPB.FoodDetailResponse_Salad{
			Salad: &asosiyPB.Salad{
				Name:        "AI salad result",
				Description: res.Steps[0],
			},
		},
	}, nil
}