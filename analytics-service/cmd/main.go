package main

import (
	"analytics-service/internal/config"
	"analytics-service/internal/handler"
	repository "analytics-service/internal/repostory"
	"analytics-service/internal/usecase"
	loadenv "analytics-service/pkg/load_env"
	"context"
	"fmt"
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

	get, err := repoCreate.GetWeeklyNutrition(ctx, 1)
if err != nil {
    log.Fatal(err)
}

fmt.Println("╔══════════════════════════════════════════════════╗")
fmt.Println("║           7-KUNLIK HAFTALIK STATISTIKA           ║")
fmt.Println("╠══════════════════════════════════════════════════╣")
fmt.Printf("║ %-12s %-10s %-10s %-10s %-6s║\n", "Kun", "Kcal", "Protein", "Fat", "Carbs")
fmt.Println("╠══════════════════════════════════════════════════╣")

for _, d := range get {
    fmt.Printf("║ %-12s %-10.1f %-10.1f %-10.1f %-6.1f║\n",
        str.TrimSpace(d.Day),
        d.Kcal,
        d.Protein,
        d.Fat,
        d.Carbs,
    )
}

fmt.Println("╚══════════════════════════════════════════════════╝")

	useCreate := usecase.NewMealUsecase(repoCreate)

	handCreeate := handler.NewHandlerConsumer(rabbitMq.Channel, useCreate)





	handCreeate.Start()

}