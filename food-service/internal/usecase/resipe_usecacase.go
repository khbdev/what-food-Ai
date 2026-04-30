package usecase

import (
	"context"
	"food-service/internal/domain"
	"food-service/internal/models"
	"errors"
)

type recipeUsecase struct {
	repo domain.RecipeRepository
}

func NewRecipeUsecase(repo domain.RecipeRepository)  {
	return &recipeUsecase{
		repo: repo,
	}
}