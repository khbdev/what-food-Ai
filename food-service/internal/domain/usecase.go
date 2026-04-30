package domain

import (
	"context"
	"food-service/internal/models"
)



type RecipeUsecase interface {
	Create(ctx context.Context, recipe *models.Recipe) error
	GetByID(ctx context.Context, id int64) (*models.Recipe, error)
	GetAll(ctx context.Context) ([]*models.Recipe, error)
	Update(ctx context.Context, recipe *models.Recipe) error
	Delete(ctx context.Context, id int64) error
}

type SaladUsecase interface {
	Create(ctx context.Context, salad *models.Salad) error
	GetByID(ctx context.Context, id int64) (*models.Salad, error)
	GetAll(ctx context.Context) ([]*models.Salad, error)
	Update(ctx context.Context, salad *models.Salad) error
	Delete(ctx context.Context, id int64) error
}
