package main

import (
	"log"
	"ymatchu_backend/route"
)

func main() {
	app := route.RouteInit()
	log.Fatal(
		app.Listen(":3000"),
	)
}
