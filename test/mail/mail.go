package main

import (
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
		"localhost:50051",
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

	res, err := client.FilterFood(nil, &asosiypb.FoodFilterRequest{
		Country:   "uzbekistan",
		MealTime:  "lunch",
		HasSalad:  true,
		KcalLimit: 700,
	})

	if err != nil {
		log.Fatal("filter error:", err)
	}

	log.Println("FILTER ITEMS:", len(res.Items))
}

// =========================
// 2 FUNCTION: DETAIL TEST
// =========================

func testFoodDetail() {

	res, err := client.GetFoodDetail(nil, &asosiypb.FoodDetailRequest{
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

	// 👉 Qaysini xohlasang chaqirasan:

	testFilterFood()
	// testFoodDetail()
}