package main

import (
	"context"
	"log"
	"user-service/internal/config"
	"user-service/internal/models"
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

	ctx, cancel := context.
user := models.User{
	Name:    "Azizbek",
	Phone:   "+998901234567",
	Age:     21,
	Address: "Tashkent, Uzbekistan",
	Email:   "azizbek@gmail.com",
	Image:   "https://example.com/avatar.png",
	Role:    models.RoleUser,
}

	repo.Create(user)

}