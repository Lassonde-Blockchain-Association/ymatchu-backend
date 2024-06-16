package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Lassonde-Blockchain-Association/ymatchu-backend/database"
	"github.com/gorilla/mux"
)

func createListingDetails(w http.ResponseWriter, r *http.Request) {
	var listingDetails database.ListingDetails
	db := database.DB
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
	database.InitDB()

	// create a new router
	r := mux.NewRouter()

	r.HandleFunc("/listings", createListingDetails).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
