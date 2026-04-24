package main

import (
	"auth-service/internal/config"
	"auth-service/internal/models"
	loadenv "auth-service/pkg/LoadEnv"
	rabbitMq "auth-service/pkg/rabbitMq"
	"log"
	"os"
)


func main(){
	loadenv.Load()

	rabbitmq := config.NewRabbit()

	_ = rabbitmq

  redis, err := config.NewRedisClient()
  if err !=nil {
	log.Fatal("Error", err)
  }

  _ = redis

  producer :=  rabbitMq.NewPublisher(rabbitmq)


 _ = 
}