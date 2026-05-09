package service

import (
	"context"
	"errors"

	"api-geteway/internal/client"

	nutritionpb ""
)

type NutritionService struct {
	nutritionClient *client.NutritionClient
}

// =========================
// INIT
// =========================

func NewNutritionService(c *client.NutritionClient) *NutritionService {
	return &NutritionService{
		nutritionClient: c,
	}
}

// =========================
// GET WEEKLY NUTRITION
// =========================

func (s *NutritionService) GetWeeklyNutrition(
	ctx context.Context,
	req *nutritionpb.WeeklyNutritionRequest,
) (*nutritionpb.WeeklyNutritionResponse, error) {

	if req.UserId == 0 {
		return nil, errors.New("user id is required")
	}

	return s.nutritionClient.GetWeeklyNutrition(req)
}