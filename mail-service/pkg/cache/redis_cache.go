package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	aipb "github.com/khbdev/what-food-proto/proto/ai"
	"github.com/redis/go-redis/v9"
)

const (
	aiMealCacheTTL    = 24 * time.Hour
	aiMealKeyPrefix   = "ai:meal:"
)

// AIMealCache - AI meal analysis natijalarini Redis da cache qiladi
// DI pattern: redis client tashqaridan inject qilinadi
type AIMealCache struct {
	rdb *redis.Client
}

func NewAIMealCache(rdb *redis.Client) *AIMealCache {
	return &AIMealCache{rdb: rdb}
}

// CachedMealAnalysis - Redis da to'liq saqlanadigan struct
// Food ma'lumotlari + AI ma'lumotlari birgalikda saqlanadi,
// cache hit bo'lganda food fetch ham, AI request ham ketmaydi
type CachedMealAnalysis struct {
	// Food ma'lumotlari
	Id           int64   `json:"id"`
	Type         string  `json:"type"`
	RestaurantId int64   `json:"restaurant_id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	ImageUrl     string  `json:"image_url"`
	VideoUrl     string  `json:"video_url"`
	Country      string  `json:"country"`
	MealTime     string  `json:"meal_time"`
	Kcal         int32   `json:"kcal"`
	Protein      float64 `json:"protein"`
	Fat          float64 `json:"fat"`
	Carbs        float64 `json:"carbs"`

	// AI ma'lumotlari
	Portion            int32              `json:"portion"`
	TotalKcal          float32            `json:"total_kcal"`
	CookingTimeMinutes int32              `json:"cooking_time_minutes"`
	Ingredients        []*aipb.Ingredient `json:"ingredients"`
	Steps              []string           `json:"steps"`
}

// buildKey - cache key yasaydi: "ai:meal:recipe:123"
func (c *AIMealCache) buildKey(foodType string, id int64) string {
	return fmt.Sprintf("%s%s:%d", aiMealKeyPrefix, foodType, id)
}

// Get - Read-through: Redis dan o'qiydi
// (nil, nil) => cache miss, (data, nil) => cache hit, (nil, err) => xato
func (c *AIMealCache) Get(ctx context.Context, foodType string, id int64) (*CachedMealAnalysis, error) {
	key := c.buildKey(foodType, id)

	val, err := c.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		// Cache miss - bu normal holat
		return nil, nil
	}
	if err != nil {
		// Redis xatosi - loglab o'tamiz, AI ga boramiz
		return nil, fmt.Errorf("redis get error: %w", err)
	}

	var result CachedMealAnalysis
	if err := json.Unmarshal([]byte(val), &result); err != nil {
		return nil, fmt.Errorf("cache unmarshal error: %w", err)
	}

	return &result, nil
}

// Set - Write-through: Redis ga yozadi, 24 soat TTL
func (c *AIMealCache) Set(ctx context.Context, foodType string, id int64, data *CachedMealAnalysis) error {
	key := c.buildKey(foodType, id)

	bytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("cache marshal error: %w", err)
	}

	if err := c.rdb.Set(ctx, key, bytes, aiMealCacheTTL).Err(); err != nil {
		return fmt.Errorf("redis set error: %w", err)
	}

	return nil
}