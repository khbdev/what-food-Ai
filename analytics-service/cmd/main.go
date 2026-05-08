package main

import (
	"analytics-service/internal/config"
	loadenv "analytics-service/pkg/load_env"
)


func main(){
	loadenv.LoadEnv()

	rabbitMq := config.NewRabbit()

	_ = rabbitMq
	db, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
}