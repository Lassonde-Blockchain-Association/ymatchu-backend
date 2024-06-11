package student

import (
	"net/http"
)

func InitializeStudentRoutes() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		http.Error(w, "Not Found", http.StatusNotFound)
	})
	mux.HandleFunc("POST /filteringRequest", filteringRequest)
	mux.HandleFunc("GET /listingDetails/{id}", listingDetails)
	return mux
}

