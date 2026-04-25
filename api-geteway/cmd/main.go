package main

import (
	"api-geteway/internal/client"
	"api-geteway/pkg/loadenv"
	"log"
	"os"
)


func main(){

	loadenv.Load()
    authUR := os.Getenv("AUTH_URL")
	authServiceClient, err := client.NewAuthClient(authUR)

	if err != nil {
		log.Fatal("Xatolik",err)
	}

	_ = authServiceClient
}