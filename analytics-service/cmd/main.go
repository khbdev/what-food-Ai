package main

import "analytics-service/internal/config"

func main(){
	rabbitMq := config.NewRabbit()

	_ = rabbitMq
}