package main

import "notifaction-service/internal/config"


func main(){

	rabbitMqConnectio := config.NewRabbit()

	_ = rabbitMqConnectio

}