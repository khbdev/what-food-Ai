package main

import (
	"log"
	"os"
)



func mustEnv(key string) string {
	v := os.Getenv(key)

	if v == "" {
		log.Fatalf("❌ %s is empty", key)
	}

	return v
}


func main(){



}