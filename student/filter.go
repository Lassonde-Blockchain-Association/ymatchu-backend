
package student
import (
	"gorm.io/gorm"
)


func filterListings(price string) func(db *gorm.DB) *gorm.DB{
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("price <= ?" , price)
	}
}

func filterLocation() func(db *gorm.DB) *gorm.DB{
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("location = ?" , "location")
	}
}


func filterUtilities(water bool, gas bool, parking int, locker bool) func(db *gorm.DB) *gorm.DB{
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("utilities = ?" , "utilities")
	}
}

func filterFeatures(rooms int, washrooms int ,  squareft int) func(db *gorm.DB) *gorm.DB{
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("features = ?" , "features")
	}
}

