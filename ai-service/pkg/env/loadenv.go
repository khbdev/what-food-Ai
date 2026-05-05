package env

import (
	"log"
	"math/big"

	"github.com/joho/godotenv"
)


func  LoadEnv(){
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error: ")
	}
}