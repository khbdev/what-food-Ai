package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

)

type Service struct {
	rdb *redis.Client
	ctx context.Context
}

func NewService(rdb *redis.Client) *Service {
	return &Service{
		rdb: rdb,
		ctx: context.Background(),
	}
}