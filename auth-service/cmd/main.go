package main

import (
	"auth-service/internal/config"
	
	loadenv "auth-service/pkg/LoadEnv"
	rabbitMq "auth-service/pkg/rabbitMq"
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

  _ = redis

  producer :=  rabbitMq.NewPublisher(rabbitmq)
 _ = producer

 userClient, err := cli

}