package route

import (
	// go fiber
	"github.com/Lassonde-Blockchain-Association/ymatchu-backend/database"
	"github.com/Lassonde-Blockchain-Association/ymatchu-backend/student"
	"github.com/gofiber/fiber/v2"
)

func RouteInit() *fiber.App {
	app := fiber.New()
	api := app.Group("/api/v1")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("YMatchu API")
	})

	// landlord
	api.Post("/landlord/createListing", func(c *fiber.Ctx) error {
		return c.SendString("Create Listing")
	})

	api.Post("/landlord/updateListing/:id", func(c *fiber.Ctx) error {
		return c.SendString("Update Listing")
	})

	api.Post("/landlord/deleteListing/:id", func(c *fiber.Ctx) error {
		return c.SendString("Delete Listing")
	})

	// student
	student := student.Student{}
	database.InitDB()
	student.DB = database.DB

	//this would filter based on parameters and return the relevant listings
	api.Post("/:studentID/filterRequest", student.Filter)

	// post listingDetails
	api.Post("/student/listingDetails/:id", student.ListingDetails)
	return app
}
