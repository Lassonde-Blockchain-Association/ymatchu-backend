package student

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Student struct {
	DB *gorm.DB
}

func (student *Student) Filter(c *fiber.Ctx) error {

	db := student.DB
	var listing Listing
	var utilities Utilities
	var features Features
	var location Location
	db.Session(&gorm.Session{NewDB: true})
	db.Scopes(filterListings(c.Query("price"))).Find(&listing)
	db.Scopes(filterLocation(c.Query("street"), c.Query("city"), c.Query("postal_code"), c.Query("country"))).Find(&location)
	parkigns, err := strconv.Atoi(c.Query("parkings"))
	if err != nil {
		parkigns = 0
	}
	db.Scopes(filterUtilities(c.Query("water") == "true", c.Query("gas") == "true", parkigns, c.Query("locker") == "true")).Find(&utilities)

	rooms, err := strconv.Atoi(c.Query("rooms"))
	if err != nil {
		rooms = 0
	}
	washrooms, err := strconv.Atoi(c.Query("washrooms"))
	if err != nil {
		washrooms = 0
	}
	squareft, err := strconv.Atoi(c.Query("squareft"))
	if err != nil {
		squareft = 0
	}
	db.Scopes(filterFeatures(rooms, washrooms, squareft)).Find(&features)

	return c.Status(200).JSON(fiber.Map{
		"listings": "Filtering Request",
	})
}

func (student *Student) ListingDetails(c *fiber.Ctx) error {
	// db := student.DB

	return c.Status(200).JSON(fiber.Map{
		"listing":   "Listing Details",
		"features":  "",
		"utilities": "",
		"location":  "",
	})
}
