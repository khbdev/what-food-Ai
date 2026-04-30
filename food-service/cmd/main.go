package main

import (
	"food-service/internal/config"
	"food-service/internal/repository"
	"food-service/pkg/env"
	"log"
)



func main(){

	env.LoadEnv()  


	postgres, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	_ = postgres

	redis, err := config.NewRedisClient()

	if err != nil {
		log.Fatal(err)
	}
   _ = redis


   resipe_salad := repository.NewRecipeRepository(postgres)

   _ = resipe_salad

}