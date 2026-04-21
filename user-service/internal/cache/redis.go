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