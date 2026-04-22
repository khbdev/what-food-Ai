package main

import (
	"log"

	"user-service/internal/config"
	"user-service/internal/usecase"

	repository "user-service/internal/repostory"
	loadenv "user-service/pkg/loadEnv"
)


func main(){
	loadenv.LoadEnv()

	sql, err := config.NewPostgresDB()
	if err != nil{
	   log.Fatal(err)
	}
	_ =  sql

	redis, err := config.NewRedisClient()
	if err != nil{
	   log.Fatal(err)
	}
	_ = redis

	repo := repository.NewUserRepository(sql)

	_ = repo

	service := usecase.

}