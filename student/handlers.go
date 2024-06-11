package student

import (
	"encoding/json"
	"net/http"
)

func filteringRequest(w http.ResponseWriter, r *http.Request){
	// get the student id from the request
	// get the preferences from the request
	// save the preferences
	// send array of matching listings
	json.NewEncoder(w).Encode("filteringRequest")
}

func listingDetails(w http.ResponseWriter, r *http.Request){
	// get the student id from the request
	// get the listing id from the request
	// get the listing details
	// return the listing details
	id := r.PathValue("id")
	json.NewEncoder(w).Encode(id)
}

