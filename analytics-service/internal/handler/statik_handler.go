package handler

import (
	"context"
	"errors"

	"analytics-service/internal/usecase"
	pb ""
)

type NutritionHandler struct {
	pb.UnimplementedNutritionServiceServer
	usecase *usecase.MealUsecase
}

// =====================
// DI (Dependency Injection)
// =====================
func NewNutritionHandler(u *usecase.MealUsecase) *NutritionHandler {
	return &NutritionHandler{
		usecase: u,
	}
}

// =====================
// RPC IMPLEMENTATION
// =====================
func (h *NutritionHandler) GetWeeklyNutrition(
	ctx context.Context,
	req *pb.WeeklyNutritionRequest,
) (*pb.WeeklyNutritionResponse, error) {

	// =====================
	// VALIDATION
	// =====================
	if req.UserId == 0 {
		return nil, errors.New("user_id is required")
	}

	// =====================
	// USECASE CALL
	// =====================
	data, err := h.usecase.GetWeeklyNutrition(ctx, uint(req.UserId))
	if err != nil {
		return nil, err
	}

	// =====================
	// MAPPING
	// =====================
	resp := &pb.WeeklyNutritionResponse{
		Days: make([]*pb.DailyNutrition, 0, len(data)),
	}

	for _, d := range data {
		resp.Days = append(resp.Days, &pb.DailyNutrition{
			Day:     d.Day,
			Kcal:    d.Kcal,
			Fat:     d.Fat,
			Carbs:   d.Carbs,
			Protein: d.Protein,
		})
	}

	return resp, nil
}