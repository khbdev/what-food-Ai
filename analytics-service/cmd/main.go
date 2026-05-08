package main

import (
	"analytics-service/internal/config"
	"analytics-service/internal/handler"
	repository "analytics-service/internal/repostory"
	"analytics-service/internal/usecase"
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

	repoCreate := repository.NewMealRepository(db)

	_ = repoCreate

	useCreate := usecase.NewMealUsecase(repoCreate)

	handCreeate := handler.NewHandlerConsumer(rabbitMq.Channel, useCreate)
}