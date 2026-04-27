package main

import (
	"log"
	"user-product-service/internal/config"
	"user-product-service/pkg/loadenv"
)



func main(){

    loadenv.LoadEnv()

	postgress, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	_ = postgress

	redis, err := con
}