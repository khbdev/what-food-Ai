package main

import "user-product-service/internal/config"



func main(){
	postgress, err := config.NewPostgresDB()
	if err != nil {
		
	}
}