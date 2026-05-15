package loadenv

import (
	"log"

	"github.com/joho/godotenv"
)



func Load(){
if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using system env")
	}
}