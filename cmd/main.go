package main

import (
	"log"
	"ymatchu-backend/route"
)

func main() {
	app := route.RouteInit()

	log.Fatal(
		app.Listen(":3000"),
	)

}
