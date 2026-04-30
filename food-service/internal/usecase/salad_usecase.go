package usecase

import (
	"context"
	"errors"

	"food-service/internal/domain"
	"food-service/internal/models"
)


type saladUsecase struct {
	repo domain.SaladRepository
}

// DI (Dependency Injection)
func NewSaladUsecase(repo domain.SaladRepository) domain.SaladUsecase {
	return &saladUsecase{
		repo: repo,
	}
}

func (u *saladUsecase) Create(ctx context.Context, salad *models.Salad) error {
	if salad.Name == "" {
		return errors.New("name is required")
	}

	return u.repo.Create(ctx, salad)
}

func (u *saladUsecase) GetByID(ctx context.Context, id int64) (*models.Salad, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	return u.repo.GetByID(ctx, id)
}

func (u *saladUsecase) GetAll(ctx context.Context) ([]*models.Salad, error) {
	return u.repo.GetAll(ctx)
}

func (u *saladUsecase) Update(ctx context.Context, salad *models.Salad) error {
	if salad.ID <= 0 {
		return errors.New("invalid id")
	}

	if salad.Name == "" {
		return errors.New("name is required")
	}

	return u.repo.Update(ctx, salad)
}

func (u *saladUsecase) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	return u.repo.Delete(ctx, id)
}