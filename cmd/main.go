package main

import (
	"fmt"
	"net/http"
	"ymatchu-backend/student"
)

func main(){
	mux := http.NewServeMux()
	studentMux := student.InitializeStudentRoutes()
	mux.Handle("/student", studentMux)
	err :=  http.ListenAndServe("localhost:8080", mux )

	if err != nil {
		fmt.Println(err.Error())
	}
}