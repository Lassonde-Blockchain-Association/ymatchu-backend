package landlord

import (
	"time"

	"github.com/Lassonde-Blockchain-Association/ymatchu-backend/database"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var validate = validator.New()

/*
*Validate Listing Struct
 */

func validateListing(listing database.Listing) error {

	validateLocationError := validateLocation(listing.Location)
	if validateLocationError != nil {
		return validateLocationError
	}

	validateUtilitiesError := validateUtilities(listing.Utilities)
	if validateUtilitiesError != nil {
		return validateUtilitiesError
	}

	validateFeaturesError := validateFeatures(listing.Features)
	if validateFeaturesError != nil {
		return validateFeaturesError
	}

	validateListingError := validate.Struct(listing)
	if validateListingError != nil {
		return validateListingError
	}
	return nil

}

func validateLocation(location database.Location) error {
	return validate.Struct(location)
}

func validateUtilities(utilities database.Utilities) error {
	return validate.Struct(utilities)
}

func validateFeatures(features database.Features) error {
	return validate.Struct(features)
}

// CreateListing creates a new listing
func CreateListing(c *fiber.Ctx) error {
	db := database.DB
	body := new(database.Listing)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	listing := database.Listing{
		LandlordID:  body.LandlordID,
		Price:       body.Price,
		Description: body.Description,
		CreatedOn:   time.Now(),
	}

	validateListingError := validateListing(*body)

	if validateListingError != nil {
		return c.Status(400).JSON(fiber.Map{
			"err": validateListingError.Error(),
			"msg": "Invalid payload",
		})
	}

	createListingResult := db.Omit(clause.Associations).Create(&listing)

	if createListingResult.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg":   "Create listing error",
			"error": createListingResult.Error,
		})
	}

	location := database.Location{
		StreetName: body.Location.StreetName,
		City:       body.Location.City,
		PostalCode: body.Location.PostalCode,
		Country:    body.Location.Country,
		ListingID:  listing.ID,
	}

	utility := database.Utilities{
		WaterIncluded:  body.Utilities.WaterIncluded,
		GasIncluded:    body.Utilities.GasIncluded,
		HydroIncluded:  body.Utilities.HydroIncluded,
		NoParkings:     body.Utilities.NoParkings,
		LockerIncluded: body.Utilities.LockerIncluded,
		ListingID:      listing.ID,
	}

	features := database.Features{
		NoRooms:     body.Features.NoRooms,
		NoWashrooms: body.Features.NoWashrooms,
		SquareFt:    body.Features.SquareFt,
		ListingID:   listing.ID,
	}

	db.Create(&utility)
	db.Create(&features)
	db.Create(&location)

	listing.Location = location
	listing.Utilities = utility
	listing.Features = features

	db.Save(&listing)
	return c.Status(200).JSON(listing)
}

// UpdateListing updates a listing
func UpdateListing(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "ID is required",
		})
	}

	var listing database.Listing
	err := db.First(&listing, "id = ?", id).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	body := new(database.Listing)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// update listing fields
	listing.LandlordID = body.LandlordID
	listing.Price = body.Price
	listing.Description = body.Description

	// update associated fields
	db.Model(&listing.Utilities).Updates(database.Utilities{
		WaterIncluded:  body.Utilities.WaterIncluded,
		GasIncluded:    body.Utilities.GasIncluded,
		HydroIncluded:  body.Utilities.HydroIncluded,
		NoParkings:     body.Utilities.NoParkings,
		LockerIncluded: body.Utilities.LockerIncluded,
	})
	db.Model(&listing.Features).Updates(database.Features{
		NoRooms:     body.Features.NoRooms,
		NoWashrooms: body.Features.NoWashrooms,
		SquareFt:    body.Features.SquareFt,
	})
	db.Model(&listing.Location).Updates(database.Location{
		StreetName: body.Location.StreetName,
		City:       body.Location.City,
		PostalCode: body.Location.PostalCode,
		Country:    body.Location.Country,
	})

	db.Save(&listing)
	return c.Status(200).JSON(listing)
}

// DeleteListing deletes a listing
func DeleteListing(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "ID is required",
		})
	}

	var listing database.Listing
	err := db.Preload(clause.Associations).First(&listing, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"error": "Listing not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&database.Utilities{}, "listing_id = ?", listing.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&database.Features{}, "listing_id = ?", listing.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&database.Location{}, "listing_id = ?", listing.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&database.PropertyMedia{}, "listing_id = ?", listing.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&listing).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Listing deleted",
	})
}

// GetListingDetails retrieves a listing's details
func GetListingDetails(c *fiber.Ctx) error {
	// retrieve the ID
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "ID is required",
		})
	}

	// fetch Listing from the database
	db := database.DB
	var listing database.Listing
	err := db.Preload(clause.Associations).First(&listing, "id = ?", id).Error

	// handle errors for not found
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"error": "Listing not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(listing)
}
