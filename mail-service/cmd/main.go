package main

import (
	"log"
	"mail-service/internal/client"
	"mail-service/pkg/loadenv"
	"os"
)



func mustEnv(key string) string {
	v := os.Getenv(key)

	if v == "" {
		log.Fatalf("❌ %s is empty", key)
	}

	return v
}


func main(){
	loadenv.Load()


		port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	aiURL := mustEnv("AI_URL")
	foodURL := mustEnv("FOOD_URL")

	aiClient := client.




}