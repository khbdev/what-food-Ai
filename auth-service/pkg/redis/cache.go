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


func (s *Service) SetOTP(otp int64, data models.RegisterRequest) error {
	key := fmt.Sprintf("otp:%d", otp)

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return s.rdb.Set(s.ctx, key, body, time.Minute).Err()
}