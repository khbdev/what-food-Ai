package main

import (
	"log"
	"user-product-service/internal/config"
)



func main(){


	
	postgress, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	_ = postgress
}