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
			i+1, item.Id, item.Type, item.Name, item., item.Country, item.MealTime,
		)
	}
}
// =========================
// 2 FUNCTION: DETAIL TEST
// =========================

func testFoodDetail() {

	ctx := context.Background()

	res, err := client.GetFoodDetail(ctx, &asosiypb.FoodDetailRequest{
		Id:      1,
		Type:    "recipe",
		Portion: 2,
	})

	if err != nil {
		log.Fatal("detail error:", err)
	}

	log.Printf("DETAIL: %+v\n", res)
}

// =========================
// MAIN (CLEAN)
// =========================

func main() {
	initClient()

	testFilterFood()
	// testFoodDetail()
}