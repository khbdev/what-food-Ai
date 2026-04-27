package usecase

import (
	"context"
	"fmt"
	"time"

	"user-product-service/internal/domain"
	"user-product-service/internal/models"
)

const categoryCacheTTL = 10 * time.Minute

type CategoryUsecase struct {
	repo  domain.CategoryRepository
	cache domain.CategoryCache
}

func NewCategoryUsecase(repo domain.CategoryRepository, cache domain.CategoryCache) *CategoryUsecase {
	return &CategoryUsecase{
		repo:  repo,
		cache: cache,
	}
}

// CREATE
func (u *CategoryUsecase) Create(ctx context.Context, req *models.Category) (*models.Category, error) {
	cat := &models.Category{
		Name: req.Name,
	}

	if err := u.repo.Create(cat); err != nil {
		return nil, fmt.Errorf("usecase.Create: %w", err)
	}

	_ = u.cache.Set(ctx, &models.CategoryWithIngredients{
		CategoryID: cat.ID,
		Name:       cat.Name,
	}, categoryCacheTTL)

	return cat, nil
}

// GET BY ID (READ THROUGH)
func (u *CategoryUsecase) GetByID(ctx context.Context, id int64) (*models.CategoryWithIngredients, error) {
	cached, err := u.cache.Get(ctx, id)
	if err == nil && cached != nil {
		return cached, nil
	}

	res, err := u.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("usecase.GetByID: %w", err)
	}

	_ = u.cache.Set(ctx, res, categoryCacheTTL)

	return res, nil
}

// GET ALL
func (u *CategoryUsecase) GetAll(ctx context.Context) ([]models.Category, error) {
	res, err := u.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("usecase.GetAll: %w", err)
	}

	return res, nil
}

// GET ALL WITH USER PRODUCTS (NO CACHE)
func (u *CategoryUsecase) GetAllWithUserProducts(ctx context.Context, userID int64) ([]models.CategoryWithIngredients, error) {
	res, err := u.repo.GetAllWithUserProducts(userID)
	if err != nil {
		return nil, fmt.Errorf("usecase.GetAllWithUserProducts: %w", err)
	}

	return res, nil
}

// UPDATE
func (u *CategoryUsecase) Update(ctx context.Context, req *models.Category) error {
	cat := &models.Category{
		ID:   req.ID,
		Name: req.Name,
	}

	if err := u.repo.Update(cat); err != nil {
		return fmt.Errorf("usecase.Update: %w", err)
	}

	_ = u.cache.Delete(ctx, cat.ID)

	updated, err := u.repo.GetByID(cat.ID)
	if err == nil {
		_ = u.cache.Set(ctx, updated, categoryCacheTTL)
	}

	return nil
}

// DELETE
func (u *CategoryUsecase) Delete(ctx context.Context, id int64) error {
	if err := u.repo.Delete(id); err != nil {
		return fmt.Errorf("usecase.Delete: %w", err)
	}

	_ = u.cache.Delete(ctx, id)

	return nil
}