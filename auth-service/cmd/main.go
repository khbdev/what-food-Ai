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


  rutingKey := os.Getenv("AUTH_ROUTING_KEY")
  
 for i := 0; i < 5; i++ {
	 producer.PublishAuthMessage(rutingKey, models.AuthMessage{
	Task: "SEND_SMS_OTP",
	Phone: "+998911013630",
	OTP: "234563",
  })
 }

  select{}
}