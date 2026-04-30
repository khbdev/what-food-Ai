package env

import "github.com/joho/godotenv"

func LoadEnv(){
	if err := godotenv.Load(); err != nil {
		
	}
}