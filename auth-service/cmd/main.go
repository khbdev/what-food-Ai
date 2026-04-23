package main

import "auth-service/internal/config"


func main(){
	
	rabbitmq := config.NewRabbit()

	_ = rabbitmq


}