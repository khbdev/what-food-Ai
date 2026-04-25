package main

import (
	"api-geteway/internal/client"
	"os"
)


func main(){
    authUR := os.Getenv("AUTH")
	authServiceClient, err := client.NewAuthClient()
}