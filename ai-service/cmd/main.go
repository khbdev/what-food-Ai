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


func main(){
	env.LoadEnv()

	groqAi := aimodel.NewGroqClient()
	
	_ = groqAi

	ctx, cencel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cencel()

reqModel := models.MealRequest{
    Name:        "Chicken Burger",
    Description: "Juicy grilled chicken burger",
    Country:     "USA",
    MealTime:    "lunch",
    Kcal:        550,
    Protein:     35,
    Fat:         20.5,
    Carbs:       60.2,
    Portion:     12,
}



	res, err := groqAi.AnalyzeMeal(ctx, reqModel)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}