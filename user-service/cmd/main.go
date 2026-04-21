package main

import (
	"log"
	"user-service/internal/config"
	loadenv "user-service/pkg/loadEnv"
)


func main(){
	loadenv.LoadEnv()

	sql, err := config.NewPostgresDB()
	if err != nil{
	   log.Fatal(err)
	}
	_ =  sql

	redis, err := config
}