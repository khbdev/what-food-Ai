package aimodel

import (
	"ai-service/internal/models"
	"context"
)

type AIModel interface {
	AnalyzeMeal(ctx context.Context, req models.MealRequest) (*models.MealResponse, error)
	AnalyzeNutrition(ctx context.Context, req models.NutritionRequest) (*models.NutritionResponse, error)
}