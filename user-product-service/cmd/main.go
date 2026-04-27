package main

import (
	"log"
	"user-product-service/internal/config"
	repository "user-product-service/internal/repostory"
	"user-product-service/pkg/loadenv"
)



func main(){

    loadenv.LoadEnv()

	postgress, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	_ = postgress

	redis, err := config.NewRedisClient()
		if err != nil {
		log.Fatal(err)
	}

	_ = redis

	repo := repository.NewCategoryRepository(postgress)
}