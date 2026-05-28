package main

import (
	"notifaction-service/internal/config"
	"notifaction-service/internal/handler"
	"notifaction-service/internal/usecase"
	"notifaction-service/pkg/loadenv"
)


func main(){

	loadenv.LoadEnv()
	rabbitMqConnectio := config.NewRabbit()

	_ = rabbitMqConnectio
   usc := usecase.NewSMSUsecase()

	handConsumer := handler.NewHandlerConsumer(rabbitMqConnectio.Channel, usc)

	handConsumer.Start()

}