package app

import (
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/customer", getAllCusters)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
