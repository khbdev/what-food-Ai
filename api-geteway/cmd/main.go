package main

import (
	"api-geteway/internal/client"
	"os"
)


func main(){
    authUR := os.Getenv("Auth")
	authServiceClient, err := client.NewAuthClient()
}