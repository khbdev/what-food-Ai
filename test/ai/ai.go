package main

import (
	pb "ai-service/proto"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAiServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// AnalyzeMeal test
	mealRes, err := client.AnalyzeMeal(ctx, &pb.MealRequest{
		Name:        "Chicken Burger",
		Description: "Juicy grilled chicken burger",
		Country:     "USA",
		MealTime:    "lunch",
		Kcal:        550,
		Protein:     35,
		Fat:         20.5,
		Carbs:       60.2,
		Portion:     12,
	})
	if err != nil {
		log.Fatalf("AnalyzeMeal error: %v", err)
	}

	fmt.Printf("Portion: %d\n", mealRes.Portion)
	fmt.Printf("Total Kcal: %.0f\n", mealRes.TotalKcal)
	fmt.Printf("Cooking Time: %d min\n", mealRes.CookingTimeMinutes)
	fmt.Println("Ingredients:")
	for _, ing := range mealRes.Ingredients {
		fmt.Printf("  - %s: %s\n", ing.Name, ing.Amount)
	}
	fmt.Println("Steps:")
	for _, step := range mealRes.Steps {
		fmt.Printf("  %s\n", step)
	}

	// AnalyzeNutrition test
	nutritionRes, err := client.AnalyzeNutrition(ctx, &pb.NutritionRequest{
		Period:     "weekly",
		AvgKcal:    2100,
		AvgProtein: 65,
		AvgFat:     80,
		AvgCarbs:   250,
	})
	if err != nil {
		log.Fatalf("AnalyzeNutrition error: %v", err)
	}

	fmt.Printf("\nLevel: %s\n", nutritionRes.Level)
	fmt.Printf("Feedback: %s\n", nutritionRes.Feedback)
}