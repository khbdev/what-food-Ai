package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"user-product-service/internal/domain"
	"user-product-service/internal/models"

	"github.com/redis/go-redis/v9"
)

type categoryCache struct {
	rdb *redis.Client
}

func NewCategoryCache(rdb *redis.Client) domain.CategoryCache {
	return &categoryCache{rdb: rdb}
}

func key(id int64) string {
	return fmt.Sprintf("category:%d", id)
}

func (c *categoryCache) Get(ctx context.Context, id int64) (*models.CategoryWithIngredients, error) {
	val, err := c.rdb.Get(ctx, key(id)).Result()
	if err != nil {
		return nil, err
	}

	var res models.CategoryWithIngredients
	if err := json.Unmarshal([]byte(val), &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *categoryCache) Set(ctx context.Context, cat *models.CategoryWithIngredients, ttl time.Duration) error {
	data, err := json.Marshal(cat)
	if err != nil {
		return err
	}

	return c.rdb.Set(ctx, key(cat.CategoryID), data, ttl).Err()
}

func (c *categoryCache) Delete(ctx context.Context, id int64) error {
	return c.rdb.Del(ctx, key(id)).Err()
}