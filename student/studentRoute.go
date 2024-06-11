package student

// set up the routes for the student

import (
	"net/http"
)

func InitializeStudentRoutes() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome to the student page"))
	})
	mux.HandleFunc("POST /preferences", preferences)
	mux.HandleFunc("GET /listings", matchedListings)
	mux.HandleFunc("GET /listings/{id}", listingDetails)
	return mux
}

