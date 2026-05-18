package usecase

import (
	"context"
	"mail-service/internal/client"
	"mail-service/internal/models"

	aipb "github.com/khbdev/what-food-proto/proto/ai"
)

// =======================
// USECASE
// =======================

type NutritionUsecase struct {
	client *client.AiClient
}

func NewNutritionUsecase(c *client.AiClient) *NutritionUsecase {
	return &NutritionUsecase{
		client: c,
	}
}

// =======================
// BUSINESS LOGIC
// =======================

func (u *NutritionUsecase) AnalyzeNutrition(
	ctx context.Context,
	req *models.NutritionRequestDTO,
) (*aipb.NutritionResponse, error) {

	var (
		totalKcal    float32
		totalFat     float32
		totalCarbs   float32
		totalProtein float32
	)

	for _, d := range req.Days {
		totalKcal += d.Kcal
		totalFat += d.Fat
		totalCarbs += d.Carbs
		totalProtein += d.Protein
	}

	n := float32(len(req.Days))
	if n == 0 {
		n = 1
	}
   pariod := "weekly"

	grpcReq := &aipb.NutritionRequest{
		Period:     pariod,
		AvgKcal:    totalKcal / n,
		AvgFat:     totalFat / n,
		AvgCarbs:   totalCarbs / n,
		AvgProtein: totalProtein / n,
	}

	// DIRECT gRPC CALL (DI orqali client)
	return u.client.AnalyzeNutrition(grpcReq)
}