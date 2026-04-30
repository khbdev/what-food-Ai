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


   resipeRepo := repository.NewRecipeRepository(postgres)

   _ = resipeRepo

   saladRepo := repository.NewSaladRepository(postgres)

   _ = saladRepo


   filterRepo := repository.NewFoodFilterRepository(postgres)

   _ = filterRepo


   restaranRepo := repository.NewRestaurantRepository(postgres)

   _ = restaranRepo



     resipe



}
