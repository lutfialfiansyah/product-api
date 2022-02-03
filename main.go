package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"product-api/config"
)

func main() {
	err := config.Routers.Run()
	if err != nil {
		log.Fatal(err)
	}
}
