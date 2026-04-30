package usecase

import (
	"context"

	"food-service/internal/domain"
	"food-service/internal/models"
)

type foodFilterUsecase struct {
	repo domain.FoodFilterRepository
}

func NewFoodFilterUsecase(r domain.FoodFilterRepository) domain.FoodFilterUsecase {
	return &foodFilterUsecase{
		repo: r,
	}
}

func (u *foodFilterUsecase) Filter(ctx context.Context, filter models.RecipeFilter) ([]models.FoodItemResponse, error) {
	return u.repo.Filter(ctx, filter)
}