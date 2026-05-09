package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	asosiypb "github.com/khbdev/what-food-proto/proto/asosiy"
)

var client asosiypb.FoodServiceClient

// =========================
// INIT CLIENT
// =========================

func initClient() {
	conn, err := grpc.Dial(
		"localhost:50057",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}

	client = asosiypb.NewFoodServiceClient(conn)
}

// =========================
// 1 FUNCTION: FILTER TEST
// =========================
func testFilterFood() {
	ctx := context.Background()

	res, err := client.FilterFood(ctx, &asosiypb.FoodFilterRequest{
		Country:   "USA",
		MealTime:  "lunch",
		HasSalad:  false,
		KcalLimit: 550,
	})

	if err != nil {
		log.Fatal("filter error:", err)
	}

	log.Println("FILTER ITEMS:", len(res.Items))

	for i, item := range res.Items {
		log.Printf("[%d] ID:%d | Type:%s | Name:%s | Kcal:%d | Country:%s | MealTime:%s\n",
			i+1, item.Id, item.Type, item.Name, item.Kcal, item.Country, item.MealTime,
		)
	}
}
// =========================
// 2 FUNCTION: DETAIL TEST
// =========================

func testFoodDetail() {
	ctx := context.Background()

	res, err := client.GetFoodDetail(ctx, &asosiypb.FoodDetailRequest{
		Id:      3,
		Type:    "recipe",
		Portion: 23,
	})

	if err != nil {
		log.Fatal("detail error:", err)
	}

	log.Println("=== FOOD INFO ===")
	log.Printf("ID: %d\n", res.Id)
	log.Printf("Type: %s\n", res.Type)
	log.Printf("Name: %s\n", res.Name)
	log.Printf("Description: %s\n", res.Description)
	log.Printf("Country: %s | MealTime: %s\n", res.Country, res.MealTime)
	log.Printf("Kcal: %d | Protein: %.1f | Fat: %.1f | Carbs: %.1f\n", res.Kcal, res.Protein, res.Fat, res.Carbs)

	log.Println("=== AI ANALYSIS ===")
	log.Printf("Portion: %d\n", res.Portion)
	log.Printf("TotalKcal: %.0f\n", res.TotalKcal)
	log.Printf("CookingTime: %d min\n", res.CookingTimeMinutes)

	log.Println("--- Ingredients ---")
	for i, ing := range res.Ingredients {
		log.Printf("[%d] %s - %s\n", i+1, ing.Name, ing.Amount)
	}

	log.Println("--- Steps ---")
	for i, step := range res.Steps {
		log.Printf("[%d] %s\n", i+1, step)
	}
}

// =========================
// MAIN (CLEAN)
// =========================

func main() {
	initClient()

	// testFilterFood()
	testFoodDetail()
}