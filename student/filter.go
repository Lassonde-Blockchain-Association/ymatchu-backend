package student

import (
	"net/http"

	"github.com/Lassonde-Blockchain-Association/ymatchu-backend/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ReqBody FilteringParams

func (student *Student) Filter(c *fiber.Ctx) error {

	db := student.DB.Session(&gorm.Session{})

	listings := []database.Listing{}
	body := new(FilteringParams)

	if err := c.BodyParser(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ReqBody = *body

	// loads all of listings limitd to 10
	err := db.Preload(clause.Associations).Limit(10).Find(&listings).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// response := []FilterResponse{}
	// // iterate through listings and append to response
	// for _, listing := range listings {
	// 	response = append(response, FilterResponse{PropertyMedia: listing.PropertyMedia, ListingID: listing.ID})
	// }

	return c.Status(http.StatusOK).JSON(fiber.Map{
		// "listings": response,
		"listing": listings,
	})
}

func FilterPrice(db *gorm.DB) *gorm.DB {
	return db.Model(&database.Listing{}).Preload("PropertyMedia").Where("price <= ?", ReqBody.Price)
}

func FilterLocation(db *gorm.DB) *gorm.DB {
	// return db.Where("city = ? or city =''", ReqBody.Location.City)
	return db.Where(
		db.Where("street_name = ?", ReqBody.Location.StreetName).Or("street_name = ''"),
	).Where(
		db.Where("city = ?", ReqBody.Location.City).Or("city = ''"),
	).Where(
		db.Where("postal_code = ?", ReqBody.Location.PostalCode).Or("postal_code = ''"),
	).Where(
		db.Where("country = ?", ReqBody.Location.Country).Or("country = ''"),
	)
}

func FilterUtilities(db *gorm.DB) *gorm.DB {
	return db.Model(&database.Utilities{}).Where("water_included = ?", ReqBody.Utility.WaterIncluded).Where("gas_included = ?", ReqBody.Utility.GasIncluded).Where("no_parkings <= ?", ReqBody.Utility.NoParkings).Where("locker_included = ?", ReqBody.Utility.LockerIncluded)
}

func FilterFeatures(db *gorm.DB) *gorm.DB {
	return db.Model(&database.Features{}).Where("no_rooms <=?", ReqBody.Features.NoRooms).Where("no_washrooms = ?", ReqBody.Features.NoWashrooms).Where("square_ft <= ?", ReqBody.Features.SquareFt)
}
