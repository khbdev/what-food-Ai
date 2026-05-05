package main

import (
	aimodel "ai-service/internal/ai-model"
	"ai-service/internal/models"
	"ai-service/pkg/env"
)


func main(){
	env.LoadEnv()

	groqAi := aimodel.NewGroqClient()
	
	_ = groqAi


	res, err := groqAi.AnalyzeMeal(models.MealRequest{
		
	})
}