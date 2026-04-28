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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default
	}

	log.Println("AUTH_URL =", authURL)
	log.Println("PORT =", port)

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


	userProduct

	log.Println("✅ User client created")
	// =========================
	// SERVICE (usecase)
	// =========================
	authService := service.NewAuthService(authClient)
	userService := service.NewUserService(userClinet)

	// =========================
	// HANDLER (HTTP)
	// =========================
	authHandler := handler.NewAuthHandler(authService)
	userHander := handler.NewUserHandler(userService)

	// =========================
	// ROUTER
	// =========================
	router := handler.SetupRouter(authHandler, userHander)

	log.Println("🚀 API Gateway running on :" + port)

	// =========================
	// START SERVER
	// =========================
	if err := router.Run(":" + port); err != nil {
		log.Fatal("❌ failed to start server:", err)
	}
}