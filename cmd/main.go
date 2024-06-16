package main

import (
	"log"
	"os"
	"github.com/Lassonde-Blockchain-Association/ymatchu-backend/database"
	"github.com/Lassonde-Blockchain-Association/ymatchu-backend/route"
)

func main() {
	app := route.RouteInit()
	database.InitDB()
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
