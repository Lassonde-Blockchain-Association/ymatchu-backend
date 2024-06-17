package student

import (
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"github.com/Lassonde-Blockchain-Association/ymatchu-backend/database"
	"strconv"
)

func (student *Student) Filter(c *fiber.Ctx) error {

	db := student.DB

	err := c.BodyParser(&student)

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"error": err.Error()},
		)
		return err
	}

	var listing database.Listing
	var utilities database.Utilities
	var features database.Features
	var location database.Location

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


func filterListings(price string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("price <= ?", price).Or("price = ?", "")
	}
}

func filterLocation(street string, city string, postal_code string, country string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("street = ? or street = ''", street).Where("city = ? or city = ''", city).Where("postal_code = ? or postal_code = ''", postal_code).Where("country = ? or country = ''", country)
	}
}

func filterUtilities(water bool, gas bool, parking int, locker bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("water_included = ?", water).Where("gas_included = ?", gas).Where("no_parkings <= ?", parking).Where("locker_included = ?", locker)
	}
}


func filterFeatures(rooms int, washrooms int, squareft int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("no_rooms <=?", rooms).Where("no_washrooms = ?", washrooms).Where("square_ft <= ?", squareft)
	}
}
