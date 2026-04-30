package domain

import (
	"context"
	
)

type FoodFilterRepository interface {
	Filter(ctx context.Context, filter models.RecipeFilter) ([]models.FoodItemResponse, error)
}