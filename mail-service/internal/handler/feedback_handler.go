package handler

import (
	"context"

	"mail-service/internal/models"
	"mail-service/internal/usecase"

	pb "github.com/khbdev/what-food-proto/proto/feedback"
)

// =======================
// GRPC HANDLER
// =======================

type NutritionHandler struct {
	pb.UnimplementedMailServiceServer
	usecase *usecase.NutritionUsecase
}

func NewNutritionHandler(u *usecase.NutritionUsecase) *NutritionHandler {
	return &NutritionHandler{
		usecase: u,
	}
}

// =======================
// RPC METHOD
// =======================

func (h *NutritionHandler) AnalyzeNutrition(
	ctx context.Context,
	req *pb.NutritionRequest,
) (*pb.NutritionResponse, error) {

	// proto → DTO
	dto := &models.NutritionRequestDTO{
		Days: make([]models.DayDTO, 0, len(req.Days)),
	}

	for _, d := range req.Days {
		dto.Days = append(dto.Days, models.DayDTO{
			Day:     d.Day,
			Kcal:    float32(d.Kcal),
			Fat:     float32(d.Fat),
			Carbs:   float32(d.Carbs),
			Protein: float32(d.Protein),
		})
	}

	// usecase call
	resp, err := h.usecase.AnalyzeNutrition(ctx, dto)
	if err != nil {
		return nil, err
	}

	// DTO → proto response
	return &pb.NutritionResponse{
		Feedback: resp.Feedback,
		Level:    resp.Level,
	}, nil
}