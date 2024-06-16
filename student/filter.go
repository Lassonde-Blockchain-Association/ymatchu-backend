package student

import (
	"gorm.io/gorm"
)

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
