package main

import (
	"auth-service/internal/config"
	loadenv "auth-service/pkg/LoadEnv"
)


func main(){
	loadenv.Load()

	rabbitmq := config.NewRabbit()

	_ = rabbitmq



}