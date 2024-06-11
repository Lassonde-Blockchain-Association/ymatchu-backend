package main

import (
	"log"
	"net/http"
	"ymatchu-backend/student"
)

func main(){
	studentMux := student.InitializeStudentRoutes()
	mux:= http.NewServeMux()
	mux.Handle("/student/", http.StripPrefix("/student", studentMux))

	// mux.Handle("/landlord/", http.StripPrefix("/landlord", landlordMux))

	log.Fatal(http.ListenAndServe(":8080", mux))
}