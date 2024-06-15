package main

import (
	"log"
	"os"
	"ymatchu_backend/route"
	"github.com/joho/godotenv"
)

func main() {
	app := route.RouteInit()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	
	log.Fatal(app.Listen(":"+ os.Getenv("PORT")))
}
