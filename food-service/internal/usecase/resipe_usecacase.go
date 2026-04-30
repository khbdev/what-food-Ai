package usecase

import (
	"context"
	"errors"

	"food-service/internal/domain"
	"food-service/internal/models"
)



type recipeUsecase struct {
	repo domain.RecipeRepository
}

func NewRecipeUsecase(repo domain.RecipeRepository) RecipeUsecase {
	return &recipeUsecase{repo: repo}
}

func (u *recipeUsecase) Create(ctx context.Context, recipe *models.Recipe) error {
	if recipe.Name == "" {
		return errors.New("name is required")
	}

	if recipe.RestaurantID <= 0 {
		return errors.New("restaurant_id is required")
	}

	return u.repo.Create(ctx, recipe)
}

func (u *recipeUsecase) GetByID(ctx context.Context, id int64) (*models.Recipe, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	return u.repo.GetByID(ctx, id)
}

func (u *recipeUsecase) GetAll(ctx context.Context) ([]*models.Recipe, error) {
	return u.repo.GetAll(ctx)
}

func (u *recipeUsecase) Update(ctx context.Context, recipe *models.Recipe) error {
	if recipe.ID <= 0 {
		return errors.New("invalid id")
	}

	if recipe.Name == "" {
		return errors.New("name is required")
	}

	return u.repo.Update(ctx, recipe)
}

func (u *recipeUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	return u.repo.Delete(ctx, id)
}