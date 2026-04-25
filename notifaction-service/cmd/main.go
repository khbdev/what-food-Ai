package main

import (
	"notifaction-service/internal/config"
	"notifaction-service/internal/handler"
	"notifaction-service/pkg/loadenv"
)


func main(){

	loadenv.Load()

	rabbitMqConnectio := config.NewRabbit()

	_ = rabbitMqConnectio
   usc := 

	handConsumer := handler.NewHandlerConsumer(rabbitMqConnectio.Channel, )

	handConsumer.Start()

}