package student

import (
	"net/http"
)


func preferences(w http.ResponseWriter, r *http.Request){
	// get the student id from the request
	// get the preferences from the request
	// save the preferences
}

func matchedListings(w http.ResponseWriter, r *http.Request){
	// get the student id from the request
	// get the matched listings
	// return the matched listings
}


func listingDetails(w http.ResponseWriter, r *http.Request){
	// get the student id from the request
	// get the listing id from the request
	// get the listing details
	// return the listing details
}

