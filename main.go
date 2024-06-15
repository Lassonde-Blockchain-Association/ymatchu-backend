package main

import (
	"time"

	uuid "github.com/google/uuid"
)

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
    MediaURL  []string  `gorm:"type:text[]"`
}