package student

import (
	"github.com/Lassonde-Blockchain-Association/ymatchu-backend/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


type ListingDetails struct {
	LandlordDetails database.LandlordDetails
	Listing         database.Listing
}

func (student *Student) ListingDetails(c *fiber.Ctx) error {
	db := student.DB
	var listings database.Listing
	var landlord database.LandlordDetails
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(&fiber.Map{"error": "id is required"})
	}

	isolatedQuery := db.Session(&gorm.Session{})
	//find the listing
	err := isolatedQuery.Preload(clause.Associations).Where("id = ?", id).First(&listings).Error
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{"error": err.Error()})
	}
	// find landlord details associated with the listing
	err2 := isolatedQuery.Omit("id", "listing").Where("id = ?", listings.LandlordID).First(&landlord).Error
	if err2 != nil {
		return c.Status(500).JSON(&fiber.Map{"error": err2.Error()})
	}

	return c.Status(200).JSON(&ListingDetails{
		LandlordDetails: landlord,
		Listing:         listings,
	})
}
