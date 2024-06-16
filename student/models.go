package student

import (
	"time"
	"github.com/google/uuid"
)

type ListingDetails struct {
	Listing     Listing
	Utilities   Utilities
	Features    Features
	Location    Location
	PropertyImages []PropertyImage
}

type LandlordDetails struct {
	ID 		uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FirstName 	string 
	MiddleName 	*string
	LastName 	string 
	Email 		string
	PhoneNumber 	string
}

type Listing struct {
    ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
    LandlordID  uuid.UUID `gorm:"type:uuid"`
    Price       int
    Description *string
    CreatedOn   time.Time
    UpdatedOn   *time.Time
    DeletedOn   *bool
}

type Utilities struct {
    ListingID    uuid.UUID `gorm:"type:uuid"`
    WaterIncluded bool
    GasIncluded   *bool
    HydroIncluded bool
    NoParkings    int
    LockerIncluded *bool
}

type Features struct {
    ListingID   uuid.UUID `gorm:"type:uuid"`
    NoRooms     int
    NoWashrooms int
    SquareFt    *float32
}

type Location struct {
    ListingID uuid.UUID `gorm:"type:uuid"`
    UnitNumber *int
    StreetName string
    City       string
    PostalCode string
    Country    string
}

type PropertyImage struct {
    ListingID uuid.UUID `gorm:"type:uuid"`
    ImageURL  []string  `gorm:"type:text[]"`
    VideoURL  []string  `gorm:"type:text[]"`
}