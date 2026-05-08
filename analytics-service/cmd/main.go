package main

import (
	"analytics-service/internal/config"
	repository "analytics-service/internal/repostory"
	loadenv "analytics-service/pkg/load_env"
	"log"
)


func main(){
	loadenv.LoadEnv()

	rabbitMq := config.NewRabbit()

	_ = rabbitMq
	db, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	_ = db

	repoCreate := repository.NewMealRepository()
}