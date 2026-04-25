package main

import "notifaction-service/internal/config"


func main(){

	load

	rabbitMqConnectio := config.NewRabbit()

	_ = rabbitMqConnectio

}