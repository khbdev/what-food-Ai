package usecase

import (
	"ai-service/internal/domain"
	"ai-service/internal/models"
	"context"
	"errors"
)

type AIUsecase struct {
	ai domain.AIModel
}

func NewAIUsecase(ai domain.AIModel) *AIUsecase {
	return &AIUsecase{ai: ai}
}

func (u *AIUsecase) AnalyzeMeal(ctx context.Context, name, description, country, mealTime string, kcal, protein, fat, carbs float32, portion int32) (*models.MealResponse, error) {
	// validation
	if name == "" {
		return nil, errors.New("name is required")
	}
	if portion <= 0 {
		return nil, errors.New("portion must be greater than 0")
	}
	if kcal <= 0 {
		return nil, errors.New("kcal must be greater than 0")
	}

	// model
	req := models.MealRequest{
		Name:        name,
		Description: description,
		Country:     country,
		MealTime:    mealTime,
		Kcal:        kcal,
		Protein:     protein,
		Fat:         fat,
		Carbs:       carbs,
		Portion:     portion,
	}

	return u.ai.AnalyzeMeal(ctx, req)
}

func (u *AIUsecase) AnalyzeNutrition(ctx context.Context, period string, avgKcal, avgProtein, avgFat, avgCarbs float32) (*models.NutritionResponse, error) {
	// validation
	if period == "" {
		return nil, errors.New("period is required")
	}
	if avgKcal <= 0 {
		return nil, errors.New("avg_kcal must be greater than 0")
	}

	// model
	req := models.NutritionRequest{
		Period:     period,
		AvgKcal:    avgKcal,
		AvgProtein: avgProtein,
		AvgFat:     avgFat,
		AvgCarbs:   avgCarbs,
	}

	return u.ai.AnalyzeNutrition(ctx, req)
}