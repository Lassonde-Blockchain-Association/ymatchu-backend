package route

import (
	// go fiber
	"github.com/Lassonde-Blockchain-Association/ymatchu-backend/database"
	"github.com/Lassonde-Blockchain-Association/ymatchu-backend/student"
	"github.com/gofiber/fiber/v2"
)

func RouteInit() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("YMatchu API")
	})
	
	// landlord
	app.Post("/createListing", func(c *fiber.Ctx) error {
		return c.SendString("Create Listing")
	})

	// student
	student := student.Student{}
	database.InitDB()
	student.DB = database.DB
	app.Post("/student/filterRequest", student.Filter)
	// post listingDetails
	app.Post("/student/listingDetails/:id", student.ListingDetails)
	return app
}
