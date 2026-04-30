package domain

import (
	"context"
	"food-service/internal/models"
)

type FoodFilterRepository interface {
	Filter(ctx context.Context, filter models.RecipeFilter) ([]models.FoodItemResponse, error)
}