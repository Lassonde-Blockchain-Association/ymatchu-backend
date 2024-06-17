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
	FirstName   string    `json:"first_name"`
	MiddleName  *string   `json:"middle_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
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
	LandlordID     string          `gorm:"type:uuid"`
	Price          int             `json:"price"`
	Description    *string         `json:"description"`
	CreatedOn      time.Time       `json:"created_on"`
	UpdatedOn      *time.Time      `json:"updated_on"`
	DeletedOn      *bool           `json:"deleted_on"`
	Utilities      Utilities       `gorm:"foreignKey:ListingID"`
	Features       Features        `gorm:"foreignKey:ListingID"`
	Location       Location        `gorm:"foreignKey:ListingID"`
	// PropertyImages []PropertyMedia `gorm:"foreignKey:ListingID"`
}

type Utilities struct {
	ListingID      string `gorm:"type:uuid;primary_key" json:"listing_id"`
	WaterIncluded  bool   `json:"water_included"`
	GasIncluded    *bool  `json:"gas_included"`
	HydroIncluded  bool   `json:"hydro_included"`
	NoParkings     int    `json:"no_parkings"`
	LockerIncluded *bool  `json:"locker_included"`
}

type Features struct {
	ListingID   string   `gorm:"type:uuid;primary_key" json:"listing_id"`
	NoRooms     int      `json:"no_rooms"`
	NoWashrooms int      `json:"no_washrooms"`
	SquareFt    *float32 `json:"square_ft"`
}

type Location struct {
	ListingID  string  `gorm:"type:uuid;primary_key" json:"listing_id"`
	UnitNumber *string `json:"unit_number"`
	StreetName string  `json:"street_name"`
	City       string  `json:"city"`
	PostalCode string  `json:"postal_code"`
	Country    string  `json:"country"`
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
