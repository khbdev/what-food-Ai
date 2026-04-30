package domain

import (
	"context"
	"food-service/internal/models"
)

type RecipeRepository interface {
	Create(ctx context.Context, recipe *models.Recipe) error

	GetByID(ctx context.Context, id int64) (*models.Recipe, error)

	GetAll(ctx context.Context) ([]*models.Recipe, error)


	Update(ctx context.Context, recipe *models.Recipe) error

	Delete(ctx context.Context, id int64) error
}

type SaladRepository interface {
	Create(ctx context.Context, salad *models.Salad) error

	GetByID(ctx context.Context, id int64) (*models.Salad, error)

	GetAll(ctx context.Context) ([]*models.Salad, error)

	Update(ctx context.Context, salad *models.Salad) error

	Delete(ctx context.Context, id int64) error
}


type RestaurantRepository interface {
	Create(ctx context.Context, r *models.Restaurant) (int64, error)
	GetByID(ctx context.Context, id int64) (*models.Restaurant, error)
	GetAll(ctx context.Context) ([]*models.Restaurant, error)
	Update(ctx context.Context, r *models.Restaurant) error
	Delete(ctx context.Context, id int64) error
}