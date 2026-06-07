package config

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

func InitRedis() *redis.Client {
	host := getEnv("REDIS_HOST", "what-food-redis")
	port := getEnv("REDIS_PORT", "6379")
	pass := getEnv("REDIS_PASSWORD", "")

	addr := fmt.Sprintf("%s:%s", host, port)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Redisga ulanishda xatolik: %v", err)
	}

	return client
}

func getEnv(key, fallback string) string {
	if val, ok := lookupEnv(key); ok {
		return val
	}
	return fallback
}

func lookupEnv(key string) (string, bool) {
	envs, _ := godotenv.Read()
	val, ok := envs[key]
	return val, ok
}
