package handler

import (
	"context"

	pb "nutrition-service/proto/nutritionpb"
	"nutrition-service/internal/usecase"
)

type NutritionHandler struct {
	pb.UnimplementedNutritionServiceServer
	usecase *usecase.NutritionUsecase
}

// =====================
// DI (Dependency Injection)
// =====================
func NewNutritionHandler(u *usecase.NutritionUsecase) *NutritionHandler {
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
	// VALIDATION (minimal handler level)
	// =====================
	if req.UserId == 0 {
		return nil, ErrInvalidUserID()
	}

	// =====================
	// USECASE CALL
	// =====================
	result, err := h.usecase.GetWeeklyNutrition(ctx, uint(req.UserId))
	if err != nil {
		return nil, err
	}

	// =====================
	// MAPPING (repo -> proto)
	// =====================
	response := &pb.WeeklyNutritionResponse{
		Days: make([]*pb.DailyNutrition, 0, len(result)),
	}

	for _, d := range result {
		response.Days = append(response.Days, &pb.DailyNutrition{
			Day:     d.Day,
			Kcal:    d.Kcal,
			Fat:     d.Fat,
			Carbs:   d.Carbs,
			Protein: d.Protein,
		})
	}

	return response, nil
}