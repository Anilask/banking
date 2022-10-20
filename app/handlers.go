package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Customer struct {
	Name    string `json:"Full_Name" xml:"Full_Name"`
	City    string `json:"City" xml:"City"`
	Zipcode string `json:"Zipcode" xml:"Zipcode"`
}

func getAllCusters(w http.ResponseWriter, r *http.Request) {
	customer := []Customer{
		{
			Name:    "Anil",
			City:    "Hyd",
			Zipcode: "505050",
		},
		{
			Name:    "Shiva",
			City:    "Khailasham",
			Zipcode: "No",
		},
	}

	if r.Header.Get("content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customer)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}

	// w.Header().Add("Content-Type", "application/xml")
	// xml.NewEncoder(w).Encode(customer)

	// w.Header().Add("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(customer)

}
