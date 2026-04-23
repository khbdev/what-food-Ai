package main

import (
	"auth-service/internal/config"
	loadenv "auth-service/pkg/LoadEnv"
	"log"
)


func main(){
	loadenv.Load()

	rabbitmq := config.NewRabbit()

	_ = rabbitmq

  redis, err := config.NewRedisClient()
  if err !=nil {
	log.Fatal("Error", err)
  }

}