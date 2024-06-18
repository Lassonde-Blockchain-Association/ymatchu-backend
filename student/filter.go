package student

import (
	"net/http"
	"github.com/Lassonde-Blockchain-Association/ymatchu-backend/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ReqBody FilteringParams

type ID struct {
	ID string `json:"id"`
}

/* 
	filterRequest handler
*/
func (student *Student) Filter(c *fiber.Ctx) error {

	db := student.DB.Session(&gorm.Session{})
	body := FilteringParams{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	ReqBody = body
	setOfIDs := NewStringSet()

	// filter by each type below

	errLocation := filterLocation(db, setOfIDs)
	if errLocation != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": errLocation.Error(),
		})
	}
	errUtilities := filterUtilities(db, setOfIDs)
	if errUtilities != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": errUtilities.Error(),
		})
	}
	errFeatures := filterFeatures(db, setOfIDs)
	if errFeatures != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": errFeatures.Error(),
		})
	}

	errPrice := filterPrice(db, setOfIDs)
	if errPrice != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": errPrice.Error(),
		})
	}

	if setOfIDs.IsEmpty() {
		return c.Status(400).JSON(fiber.Map{
			"error": "No listings found",
		})
	}
	// filter by each type above 

	// Iterate through the set of IDs and fetch the listings
	listings := []database.Listing{}
	ids := setOfIDs.Elements()
	for _, id := range ids {
		tempListing := database.Listing{}
		err := db.Preload(clause.Associations).Where("id = ?", id).Find(&tempListing).Error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		listings = append(listings, tempListing)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"listing": listings,
	})
}

func filterPrice(db *gorm.DB, set StringSet) error {
	ids := []ID{}
	err := db.Model(&database.Listing{}).Select("id").Where("price <= ?", ReqBody.Price).Find(&ids).Error
	if err != nil {
		return err
	}

	for _, id := range ids {
		set.Add(id.ID)
	}

	return err
}

func filterLocation(db *gorm.DB, set StringSet) error {
	ids := []ID{}
	err := db.Model(&database.Location{}).Select("listing_id as id").Where("city = ?", ReqBody.Location.City).Where("street_name = ?", ReqBody.Location.StreetName).Where("postal_code = ?", ReqBody.Location.PostalCode).Where("country = ?", ReqBody.Location.Country).Find(&ids).Error
	if err != nil {
		return err
	}
	for _, id := range ids {
		set.Add(id.ID)
	}
	return err
}

func filterUtilities(db *gorm.DB, set StringSet) error {
	ids := []ID{}
	err := db.Model(&database.Utilities{}).Select("listing_id as id").Where("water_included = ? or water_included = false", ReqBody.Utility.WaterIncluded).Where("gas_included = ? or gas_included = false", ReqBody.Utility.GasIncluded).Where("no_parkings <= ? or no_parkings = 0", ReqBody.Utility.NoParkings).Where("locker_included = ? or locker_included = false", ReqBody.Utility.LockerIncluded).Find(&ids).Error
	if err != nil {
		return err
	}
	for _, id := range ids {
		set.Add(id.ID)
	}
	return err
}

func filterFeatures(db *gorm.DB, set StringSet) error {
	ids := []ID{}
	err := db.Model(&database.Features{}).Select("listing_id as id").Where("no_rooms <=? or no_rooms = 0", ReqBody.Features.NoRooms).Where("no_washrooms = ?  or no_washrooms = 0", ReqBody.Features.NoWashrooms).Where("square_ft <= ? or square_ft = 0", ReqBody.Features.SquareFt).Find(&ids).Error
	if err != nil {
		return err
	}
	for _, id := range ids {
		set.Add(id.ID)
	}
	return err
}
