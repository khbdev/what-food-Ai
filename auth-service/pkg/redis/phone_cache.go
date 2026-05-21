package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type PhoneCache struct {
	client *redis.Client
	ctx    context.Context  // ✅ ctx qo'shildi
}

func NewPhoneCache(client *redis.Client) *PhoneCache {
	return &PhoneCache{
		client: client,
		ctx:    context.Background(),  // ✅ initsializatsiya
	}
}

const userPhoneTTL = time.Minute

func userPhoneKey(phone string) string {
	return fmt.Sprintf("user_phone:%s", phone)
}

// Get
func (r *PhoneCache) Get(phone string) (bool, error) {
	key := userPhoneKey(phone)

	_, err := r.client.Get(r.ctx, key).Result()  // ✅ r.ctx
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

// Set
func (r *PhoneCache) Set(phone string) error {
	key := userPhoneKey(phone)

	return r.client.Set(
		r.ctx,  // ✅ r.ctx
		key,
		"1",
		userPhoneTTL,
	).Err()
}