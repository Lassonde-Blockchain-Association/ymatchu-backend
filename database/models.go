package database

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type Base struct {
	ID string `gorm:"type:uuid;primary_key" json:"id"`
}

type LandlordDetails struct {
	Base
	FirstName   string    `json:"first_name" validate:"required, min=1"`
	MiddleName  *string   `json:"middle_name"`
	LastName    string    `json:"last_name" validate:"required, min=1"`
	Email       string    `json:"email" validate:"required, email"`
	PhoneNumber string    `json:"phone_number" valiedate:"required"`
	Listing     []Listing `gorm:"foreignKey:LandlordID"`
}
type StudentDetails struct {
	Base
	FirstName   string  `json:"first_name"`
	MiddleName  *string `json:"middle_name"`
	LastName    string  `json:"last_name"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phone_number"`
}

type Listing struct {
	Base
	LandlordID  string     `gorm:"foreignKey:LandlordID"`
	Price       int        `json:"price" validate:"required"`
	Description *string    `json:"description"`
	CreatedOn   time.Time  `json:"created_on"`
	UpdatedOn   *time.Time `json:"updated_on"`
	DeletedOn   *bool      `json:"deleted_on"`
	Utilities   Utilities  `gorm:"foreignKey:ListingID"`
	Features    Features   `gorm:"foreignKey:ListingID"`
	Location    Location   `gorm:"foreignKey:ListingID"`
	// PropertyImages []PropertyMedia `gorm:"foreignKey:ListingID"`
}

type Utilities struct {
	ListingID      string `gorm:"type:uuid;primary_key" json:"listing_id" `
	WaterIncluded  bool   `json:"water_included" validate:"required"`
	GasIncluded    *bool  `json:"gas_included"`
	HydroIncluded  bool   `json:"hydro_included" validate:"required"`
	NoParkings     int    `json:"no_parkings" validate:"required,min=0"`
	LockerIncluded *bool  `json:"locker_included"`
}

type Features struct {
	ListingID   string   `gorm:"type:uuid;primary_key" json:"listing_id" `
	NoRooms     int      `json:"no_rooms" validate:"required"`
	NoWashrooms int      `json:"no_washrooms" validate:"required"`
	SquareFt    *float32 `json:"square_ft"`
}

type Location struct {
	ListingID  string  `gorm:"type:uuid;primary_key" json:"listing_id" `
	UnitNumber *string `json:"unit_number"`
	StreetName string  `json:"street_name" validate:"required"`
	City       string  `json:"city" validate:"required"`
	PostalCode string  `json:"postal_code" validate:"required"`
	Country    string  `json:"country" validate:"required"`
}

type PropertyMedia struct {
	ListingID string   `gorm:"type:uuid;primary_key"`
	ImageURL  []string `gorm:"type:text[]" json:"image_url"`
	VideoURL  []string `gorm:"type:text[]" json:"video_url"`
}

// type ListingDetails struct {
// 	Listing        Listing
// 	Utilities      Utilities
// 	Features       Features
// 	Location       Location
// 	PropertyImages []PropertyMedia
// }

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
