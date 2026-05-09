package main

import (
	"log"
	"os"

	"api-geteway/internal/client"
	"api-geteway/internal/handler"
	"api-geteway/internal/service"
	"api-geteway/pkg/loadenv"
)

func main() {

	// =========================
	// LOAD ENV
	// =========================
	loadenv.Load()

	authURL := os.Getenv("AUTH_URL")
	if authURL == "" {
		log.Fatal("❌ AUTH_URL is empty")
	}
userURL := os.Getenv("USER_URL")
	if authURL == "" {
		log.Fatal("❌ AUTH_URL is empty")
	}

	userProductUrl := os.Getenv("USERPRODUCT_URL")
	if authURL == "" {
		log.Fatal("❌ USERPRODUCT_URL is empty")
	}
	foodUrl := os.Getenv("FOOD_URL")
	if foodUrl == "" {
			log.Fatal("❌ FOOD_URL is empty")
	}
	statikUrl := os.Getenv("STATIK_URL")
if statikUrl == "" {
	log.Fatal("Statik not fount")
}


	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default
	}



	// =========================
	// CLIENT (gRPC)
	// =========================
	authClient, err := client.NewAuthClient(authURL)
	
	if err != nil {
		log.Fatal("❌ Failed to connect auth service:", err)
	}
	defer authClient.Close()



	log.Println("✅ Auth client created")

	userClinet, err := client.NewUserClient(userURL)

		if err != nil {
		log.Fatal("❌ Failed to connect user service:", err)
	}
	defer authClient.Close()
log.Println("✅ User client created")

	userProductClient, err := client.NewUserProductClient(userProductUrl)

		if err != nil {
		log.Fatal("❌ Failed to connect user product	 service:", err)
	}
	defer userProductClient.Close()
  log.Println("✅ user Product client created")
		

	foodClient, err := client.NewFoodClient(foodUrl)
	if err != nil {
			log.Fatal("❌ Failed to connect food service:", err)
	}

	_ = foodClient
log.Println("✅ Food client created")

statikClient, err := client.NewNutritionClient(statikUrl)
if  err != nil
	// =========================
	// SERVICE (usecase)
	// =========================
	authService := service.NewAuthService(authClient)
	userService := service.NewUserService(userClinet)
	userCategory := service.NewCategoryService(userProductClient)
	userProduct := service.NewProductService(userProductClient)
	foodService := service.NewFoodService(foodClient)


	// =========================
	// HANDLER (HTTP)
	// =========================
	authHandler := handler.NewAuthHandler(authService)
	userHander := handler.NewUserHandler(userService)
	categoryHandler := handler.NewCategoryHandler(userCategory)
	productHandler := handler.NewIngredientHandler(userProduct)
	foodHandler := handler.NewFoodHandler(foodService)


	// =========================
	// ROUTER
	// =========================
	router := handler.SetupRouter(authHandler, userHander, categoryHandler, productHandler, foodHandler)

	log.Println("🚀 API Gateway running on :" + port)

	// =========================
	// START SERVER
	// =========================
	if err := router.Run(":" + port); err != nil {
		log.Fatal("❌ failed to start server:", err)
	}
}