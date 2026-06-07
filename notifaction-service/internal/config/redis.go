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
	addr := getEnv("REDIS_HOST", "localhost:6379")
	pass := getEnv("REDIS_PASSWORD", "")

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Redisga ulanishda xatolik: %v", err)
	}

	fmt.Println("Redis ulandi:", addr)

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
