package main

import (
	"api-geteway/internal/client"
	"os"
)


func main(){
    authUR := os.Getenv("AUTH_URL")
	authServiceClient, err := client.NewAuthClient(authUR)

	if err ! {
		
	}
}