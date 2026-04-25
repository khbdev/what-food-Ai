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

	log.Println("AUTH_URL =", authURL)

	// =========================
	// CLIENT (gRPC)
	// =========================
	authClient, err := client.NewAuthClient(authURL)
	if err != nil {
		log.Fatal("❌ Failed to connect auth service:", err)
	}
	defer authClient.Close()

	log.Println("✅ Auth client created")

	// =========================
	// SERVICE (usecase)
	// =========================
	authService := service.NewAuthService(authClient)

	// =========================
	// HANDLER (HTTP)
	// =========================
	authHandler := handler.NewAuthHandler(authService)

	// =========================
	// ROUTER
	// =========================
	router := handler.SetupRouter(authHandler)

	log.Println("🚀 API Gateway running on :8085")

	// 🔥 START SERVER (MUHIM)
	if err := router.Run(":8080"); err != nil {
		log.Fatal("❌ failed to start server:", err)
	}
}