package main

import (
	"api-geteway/internal/client"
	"log"
	"os"
)


func main(){
    authUR := os.Getenv("AUTH_URL")
	authServiceClient, err := client.NewAuthClient(authUR)

	if err != nil {
		log.Fatal("Xatolik",err)
	}

	_ = authServiceClient
}