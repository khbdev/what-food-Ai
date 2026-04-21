package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"user-service/internal/domain"
	"user-service/internal/models"

	"github.com/redis/go-redis/v9"
)

type userCache struct {
	rdb *redis.Client
}

func NewUserCache(rdb *redis.Client) domain.UserCache {
	return &userCache{
		rdb: rdb,
	}
}

func key(id uint) string {
	return fmt.Sprintf("user:%d", id)
}



func (c *userCache) GetUser(ctx context.Context, id uint) (*models.User, error) {
	val, err := c.rdb.Get(ctx, key(id)).Result()
	if err != nil {
		return nil, err
	}

	var user models.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return nil, err
	}

	return &user, nil
}
// SET

func (c *userCache) SetUser(ctx context.Context, user *models.User, ttl time.Duration) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return c.rdb.Set(ctx, key(user.ID), data, ttl).Err()
}


// Delete

func (c *userCache) DeleteUser(ctx context.Context, id uint) error {
	return c.rdb.Del(ctx, key(id)).Err()
}