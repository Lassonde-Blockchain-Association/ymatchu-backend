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

	app.Post("/updateListing/:id", func(c *fiber.Ctx) error {
		return c.SendString("Update Listing")
	})

	app.Post("/deleteListing/:id", func(c *fiber.Ctx) error {
		return c.SendString("Delete Listing")
	})

	// student
	student := student.Student{}
	database.InitDB()
	student.DB = database.DB

	//this would filter based on parameters and return the relevant listings
	app.Post("/:studentID/filterRequest", student.Filter)

	// post listingDetails
	app.Post("/student/listingDetails/:id", student.ListingDetails)
	return app
}
