package main

import (
	aimodel "ai-service/internal/ai-model"
	"ai-service/internal/models"
	"ai-service/pkg/env"
	"context"
	"fmt"
	"log"
	"time"
)


func main() {
	env.LoadEnv()

	groqAi := aimodel.NewGroqClient()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	reqModel := models.NutritionRequest{
		Period:     "weekly",
		AvgKcal:    2100,
		AvgProtein: 65,
		AvgFat:     80,
		AvgCarbs:   250,
	}

	res, err := groqAi.AnalyzeNutrition(ctx, reqModel)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Level:    %s\n", res.Level)
	fmt.Printf("Feedback: %s\n", res.Feedback)
}