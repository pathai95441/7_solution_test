package main

import (
	"fmt"
	"log"
	"net/http"

	"pie_fire_dire_rest/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/beef/summary", handler.BeefSummaryHandler).Methods("GET")

	http.Handle("/", r)
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
