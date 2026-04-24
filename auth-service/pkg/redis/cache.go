package redis

import (
	"auth-service/internal/models"
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


func (s *Service) SetOTP(otp int64, data models.RegisterRequest) error {
	key := fmt.Sprintf("otp:%d", otp)

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return s.rdb.Set(s.ctx, key, body, time.Minute).Err()
}

func (s *Service) GetOTP(otp int64) (*models.RegisterRequest, error) {
	key := fmt.Sprintf("otp:%d", otp)

	val, err := s.rdb.Get(s.ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var data models.RegisterRequest
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		return nil, err
	}

	_ = s.rdb.Del(s.ctx, key).Err()

	return &data, nil
}