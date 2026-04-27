package domain

import (
	"context"
	"time"

	"user-product-service/internal/models"
)

type CategoryCache interface {
	Get(ctx context.Context, id int64) (*models.CategoryWithIngredients, error)
	Set(ctx context.Context, c *models.CategoryWithIngredients, ttl time.Duration) error
	Delete(ctx context.Context, id int64) error
}