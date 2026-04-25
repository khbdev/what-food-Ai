package main

import (
	"notifaction-service/internal/config"
	"notifaction-service/pkg/loadenv"
)


func main(){

	loadenv.Load()

	rabbitMqConnectio := config.NewRabbit()

	_ = rabbitMqConnectio

}