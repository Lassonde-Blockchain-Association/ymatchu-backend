package route

import (
	// go fiber
	"github.com/Lassonde-Blockchain-Association/ymatchu-backend/database"
	"github.com/Lassonde-Blockchain-Association/ymatchu-backend/landlord"
	"github.com/Lassonde-Blockchain-Association/ymatchu-backend/student"
	"github.com/gofiber/fiber/v2"
)

func RouteInit() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("YMatchu API")
	})

	// landlord
	// landlord := landlord.Landlord{}
	database.InitDB()
	// landlord.DB = database.DB
	app.Post("/landlord/createListing", landlord.CreateListing)
	app.Get("/landlord/listing/:id", landlord.GetListingDetails)
	app.Post("/landlord/updateListing/:id", landlord.UpdateListing)
	app.Post("/landlord/deleteListing/:id", landlord.DeleteListing)
	// app.Post("/landlord/createListing", func(c *fiber.Ctx) error {
	// 	return c.SendString("Create Listing")
	// })

	// app.Post("/landlord/updateListing/:id", func(c *fiber.Ctx) error {
	// 	return c.SendString("Update Listing")
	// })

	// app.Post("/landlord/deleteListing/:id", func(c *fiber.Ctx) error {
	// 	return c.SendString("Delete Listing")
	// })

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
