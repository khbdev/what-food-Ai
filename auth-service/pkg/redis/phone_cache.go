package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type PhoneCache struct {
	client *redis.Client
}

func NewRedis(client *redis.Client) *PhoneCache {
	return &PhoneCache{
		client: client,
	}
}

const userPhoneTTL = time.Minute

func userPhoneKey(phone string) string {
	return fmt.Sprintf("user_phone:%s", phone)
}

// Get
// redisda phone mavjud bo‘lsa true qaytaradi
func (r *PhoneCache) Get(ctx context.Context, phone string) (bool, error) {
	key := userPhoneKey(phone)

	_, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

// Set
// phone numberni redisga 1 minut TTL bilan saqlaydi
func (r *PhoneCache) Set(ctx context.Context, phone string) error {
	key := userPhoneKey(phone)

	return r.client.Set(
		ctx,
		key,
		1,
		userPhoneTTL,
	).Err()
}