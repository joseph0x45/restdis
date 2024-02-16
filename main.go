package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	log.Println("Hello nigga")
}
