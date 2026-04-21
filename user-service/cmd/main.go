package main

import (
	"log"
	loadenv "user-service/pkg/loadEnv"
)


func main(){
	loadenv.LoadEnv()

	sql, err := con.NewPostgresDB()
	if err != nil{
	   log.Fatal(err)
	}

	_ = sql
}