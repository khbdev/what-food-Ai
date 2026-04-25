package main

import (
	"api-geteway/internal/client"
	"api-geteway/pkg/loadenv"
	"log"
	"os"
)

func main() {
	loadenv.Load()

	authURL := os.Getenv("AUTH_URL")
	if authURL == "" {
		log.Fatal("❌ AUTH_URL is empty")
	}

	log.Println("AUTH_URL =", authURL)

	authServiceClient, err := client.NewAuthClient(authURL)
	if err != nil {
		log.Fatal("❌ Failed to connect auth service:", err)
	}

	log.Println("✅ Auth client created")

	_ = authServiceClient
}