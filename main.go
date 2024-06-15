package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
    ImageURL  []string  `gorm:"type:text[]"`
    VideoURL  []string  `gorm:"type:text[]"`
}

var db *gorm.DB


func initDB() {
    // loading environment variables 
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }

    // build the DSN from .env 
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )

    var err error
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    db.AutoMigrate(&LandlordDetails{}, &Listing{}, &Utilities{}, &Features{}, &Location{}, &PropertyImage{})
    log.Println("Database migrated successfully!")
}


type ListingDetails struct {
	Listing     Listing
	Utilities   Utilities
	Features    Features
	Location    Location
	PropertyImages []PropertyImage
}

func createListingDetails(w http.ResponseWriter, r *http.Request) {
	var listingDetails ListingDetails
	if err := json.NewDecoder(r.Body).Decode(&listingDetails); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the Listing
	if err := db.Create(&listingDetails.Listing).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the ListingID for the related entities
	listingID := listingDetails.Listing.ID
	listingDetails.Utilities.ListingID = listingID
	listingDetails.Features.ListingID = listingID
	listingDetails.Location.ListingID = listingID
	for i := range listingDetails.PropertyImages {
		listingDetails.PropertyImages[i].ListingID = listingID
	}

	// Create related entities
	if err := db.Create(&listingDetails.Utilities).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := db.Create(&listingDetails.Features).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := db.Create(&listingDetails.Location).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := db.Create(&listingDetails.PropertyImages).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(listingDetails)
}

func main() {
    // initialization
    initDB()

    // create a new router
    r := mux.NewRouter()

	r.HandleFunc("/listings", createListingDetails).Methods("POST")

    log.Fatal(http.ListenAndServe(":8000", r))
}