package main

import (
	"food-service/internal/config"
	"food-service/pkg/env"
)



func main(){

	env.LoadEnv()  


	postgres, err := config.NewPostgresDB()
	if err !=  {
		
	}

}