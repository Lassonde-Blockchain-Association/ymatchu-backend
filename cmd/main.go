package main

import (
	"fmt"
	"net/http"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Hello World")
	})
	err :=  http.ListenAndServe("localhost:8080", mux)

	if err != nil {
		fmt.Println(err.Error())
	}
}