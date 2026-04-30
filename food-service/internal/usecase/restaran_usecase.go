package usecase

import (
	"context"
	"errors"
	"time"

	"food-service/internal/domain"
	"food-service/internal/models"
)

type restaurantUsecase struct {
	repo domain.RestaurantRepository
}

// DI (Dependency Injection)
func NewRestaurantUsecase(repo domain.RestaurantRepository) domain.RestaurantUsecase {
	return &restaurantUsecase{
		repo: repo,
	}
}

func (u *restaurantUsecase) Create(ctx context.Context, r *models.Restaurant) (int64, error) {
	if r.RestaurantName == "" {
		return 0, errors.New("restaurant_name is required")
	}

	if r.Description == "" {
		return 0, errors.New("description is required")
	}

	// business logic: created_at set usecase layerda
	r.CreatedAt = time.Now()

	return u.repo.Create(ctx, r)
}

func (u *restaurantUsecase) GetByID(ctx context.Context, id int64) (*models.Restaurant, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	return u.repo.GetByID(ctx, id)
}

func (u *restaurantUsecase) GetAll(ctx context.Context) ([]*models.Restaurant, error) {
	return u.repo.GetAll(ctx)
}

func (u *restaurantUsecase) Update(ctx context.Context, r *models.Restaurant) error {
	if r.ID <= 0 {
		return errors.New("invalid id")
	}

	if r.RestaurantName == "" {
		return errors.New("restaurant_name is required")
	}

	return u.repo.Update(ctx, r)
}

func (u *restaurantUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	return u.repo.Delete(ctx, id)
}