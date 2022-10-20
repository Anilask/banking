package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// 	http.HandleFunc("/customer", getAllCusters)

	// 	log.Fatal(http.ListenAndServe("localhost:8080", nil))
	//mux := http.NewServeMux()
	mux := mux.NewRouter()
	// define routes
	mux.HandleFunc("/customer", getAllCusters)

	// Start the server

	log.Fatal(http.ListenAndServe("localhost:8080", mux))

}
