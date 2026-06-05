package main

import (
	"log"
	"os"

	"api-geteway/internal/client"
	"api-geteway/internal/client/mailservice"
	"api-geteway/internal/handler"
	mailhandler "api-geteway/internal/handler/mail_handler"
	"api-geteway/internal/service"
	"api-geteway/internal/service/mail"
	"api-geteway/pkg/loadenv"
)

func mustEnv(key string) string {
	v := os.Getenv(key)

	if v == "" {
		log.Fatalf("❌ %s is empty", key)
	}

	return v
}

func main() {

	// =========================
	// LOAD ENV
	// =========================

	loadenv.LoadEnv()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// =========================
	// ENV
	// =========================

	authURL := mustEnv("AUTH_URL")
	userURL := mustEnv("USER_URL")
	userProductURL := mustEnv("USERPRODUCT_URL")
	foodURL := mustEnv("FOOD_URL")
	statikURL := mustEnv("STATIK_URL")
	mailURL := mustEnv("MAIL_URL")
	feedBack := mustEnv("MAIL_URL")


	// =========================
	// gRPC CLIENTS
	// =========================

	authClient, err := client.NewAuthClient(authURL)
	if err != nil {
		log.Fatal(err)
	}
	defer authClient.Close()

	userClient, err := client.NewUserClient(userURL)
	if err != nil {
		log.Fatal(err)
	}
	defer userClient.Close()

	userProductClient, err := client.NewUserProductClient(userProductURL)
	if err != nil {
		log.Fatal(err)
	}
	defer userProductClient.Close()

	foodClient, err := client.NewFoodClient(foodURL)
	if err != nil {
		log.Fatal(err)
	}

	statikClient, err := client.NewNutritionClient(statikURL)
	if err != nil {
		log.Fatal(err)
	}
	defer statikClient.Close()

	mailClient, err := mailservice.NewFoodClient(mailURL)
	if err != nil {
		log.Fatal(err)
	}
	defer mailClient.Close()

	feedMailClient, err := client.NewFeedbackClient(feedBack)
	if err != nil {
		log.Fatal(err)
	}
	userDashboard_UserService_client, err := client.NewDashboardClient(userURL)
	if err != nil {
		log.Fatal(err)
	}
	defer feedMailClient.Close()

	log.Println("✅ All gRPC clients connected")

	// =========================
	// SERVICES
	// =========================

	authService := service.NewAuthService(authClient)
	userService := service.NewUserService(userClient)
	categoryService := service.NewCategoryService(userProductClient)
	productService := service.NewProductService(userProductClient)
	foodService := service.NewFoodService(foodClient)
	statikService := service.NewNutritionService(statikClient)
	mailService := mail.NewFoodService(mailClient)
	feedbackService := service.NewFeedbackService(feedMailClient)
	userDashboardService := service.NewDashboardService(userDashboard_UserService_client)


	// =========================
	// HANDLERS
	// =========================

	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	productHandler := handler.NewIngredientHandler(productService)
	foodHandler := handler.NewFoodHandler(foodService)
	statikHandler := handler.NewNutritionHandler(statikService)
	mailHand := mailhandler.NewFoodHandler(mailService)
	feedHand := handler.NewFeedbackHandler(feedbackService)
	userDashboardHandler := handler.NewDashboardHandler(userDashboardService)

	// =========================
	// ROUTER
	// =========================
router := handler.SetupRouter(
	authHandler,
	userHandler,
	categoryHandler,
	productHandler,
	foodHandler,
	mailHand,
	statikHandler,
	feedHand,
	userDashboardHandler,
)

	// =========================
	// START SERVER
	// =========================

	log.Println("🚀 API Gateway running on :" + port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}