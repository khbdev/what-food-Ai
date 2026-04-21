package main

import (
	"log"
	
)


func main(){
	.LoadEnv()

	sql, err := config.NewPostgresDB()
	if err != nil{
	   log.Fatal(err)
	}

	_ = sql
}