package main

import (
	"analytics-service/internal/config"
	"analytics-service/internal/handler"
	repository "analytics-service/internal/repostory"
	"analytics-service/internal/usecase"
	loadenv "analytics-service/pkg/load_env"
	"context"
	"log"
	"time"
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

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	get, err := repoCreate.GetWeeklyNutrition(ctx, 1)

	useCreate := usecase.NewMealUsecase(repoCreate)

	handCreeate := handler.NewHandlerConsumer(rabbitMq.Channel, useCreate)





	handCreeate.Start()

}