package main

import (

	"user-service/user-service/internal/config"
	loadenv "user-service/user-service/pkg/loadEnv"
)


func main(){
	loadenv.LoadEnv()

	sql, err := config.NewPostgresDB()
	if err != nil{
	
	}
}