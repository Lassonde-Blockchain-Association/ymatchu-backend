package student

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID string `gorm:"type:uuid;primary_key"`
}

type LandlordDetails struct {
	// ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	Base
	FirstName   string    `json:"first_name"`
	MiddleName  *string   `json:"middle_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Listings    []Listing `gorm:"foreignKey:LandlordID;references:ID"`
}

type Listing struct {
	// ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	Base
	LandlordID     uuid.UUID       `gorm:"type:uuid"`
	Price          int             `json:"price"`
	Description    *string         `json:"description"`
	CreatedOn      time.Time       `json:"created_on"`
	UpdatedOn      *time.Time      `json:"updated_on"`
	DeletedOn      *bool           `json:"deleted_on"`
	Utilities      Utilities       `json:"utilities"`
	Features       Features        `json:"features"`
	Location       Location        `json:"location"`
	// PropertyImages []PropertyImage `json:"property_images"`
}

type Utilities struct {
	ListingID      uuid.UUID `gorm:"type:uuid"`
	WaterIncluded  bool      `json:"water_included"`
	GasIncluded    *bool     `json:"gas_included"`
	HydroIncluded  bool      `json:"hydro_included"`
	NoParkings     int       `json:"no_parkings"`
	LockerIncluded *bool     `json:"locker_included"`
}

type Features struct {
	ListingID   uuid.UUID `gorm:"type:uuid"`
	NoRooms     int       `json:"no_rooms"`
	NoWashrooms int       `json:"no_washrooms"`
	SquareFt    *float32  `json:"square_ft"`
}

type Location struct {
	ListingID  uuid.UUID `gorm:"type:uuid"`
	UnitNumber *int
	StreetName string `json:"street_name"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

type PropertyImage struct {
	ListingID uuid.UUID `gorm:"type:uuid"`
	ImageURL  []string  `gorm:"type:text[]"`
	VideoURL  []string  `gorm:"type:text[]"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
