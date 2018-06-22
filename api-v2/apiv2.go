package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"io"
	"encoding/json"
	"math/rand"
	"strconv"
)

//property struct

type Property struct {
	ID       string `json:"id,omitempty"`
	Type     string `json:"type,omitempty"`
	Bedrooms int    `json:"bedrooms,omitempty"`
	Price    int    `json:"price,omitempty"`
	Area     *Area  `json:"area,omitempty"`

}

type Area struct {
	Street 	 string   `json:"street,omitempty"`
	City   	 string   `json:"city,omitempty"`
	Province string   `json:"province,omitempty"`
}

var properties []Property


func healthTestEndpoint(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func getProperties(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(properties)
}

func getProperty(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range properties {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createProperty(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var property Property
	_ = json.NewDecoder(r.Body).Decode(&property)
	property.ID = strconv.Itoa(rand.Intn(10000000)) //mock ID - remove for production use
	properties = append(properties, property)
	json.NewEncoder(w).Encode(property)
}

func updateProperty(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range properties {
		if item.ID == params["id"]{
			properties = append(properties[:index], properties[index+1:]...)
			var property Property
			_ = json.NewDecoder(r.Body).Decode(&property)
			property.ID = strconv.Itoa(rand.Intn(10000000))
			properties = append(properties, property)
			json.NewEncoder(w).Encode(property)
			return
		}
	}
}

func deleteProperty(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range properties {
		if item.ID == params["id"] {
			properties = append(properties[:index], properties[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(properties)
}


func main(){

	r := mux.NewRouter()

	// Mock Data -TODO - implement database

	properties = append(properties, Property{ID: "1", Type: "House", Bedrooms: 6, Price: 559599, Area: &Area{Street: "Sunset boulevard", City: "Los Angeles", Province: "Los Angeles County"}})
	properties = append(properties, Property{ID: "2", Type: "House", Bedrooms: 4, Price: 450000, Area: &Area{Street: "Sunset boulevard", City: "Los Angeles", Province: "Los Angeles County"}})
	properties = append(properties, Property{ID: "3", Type: "House", Bedrooms: 4, Price: 400000, Area: &Area{Street: "Sunset boulevard", City: "Los Angeles", Province: "Los Angeles County"}})
	properties = append(properties, Property{ID: "4", Type: "Apartment", Bedrooms: 4, Price: 225000, Area: &Area{Street: "Montana Avenue", City: "Los Angeles", Province: "Los Angeles County"}})
	properties = append(properties, Property{ID: "5", Type: "Apartment", Bedrooms: 4, Price: 235000, Area: &Area{Street: "Montana Avenue", City: "Los Angeles", Province: "Los Angeles County"}})
	properties = append(properties, Property{ID: "6", Type: "Apartment", Bedrooms: 4, Price: 245000, Area: &Area{Street: "Montana Avenue", City: "Los Angeles", Province: "Los Angeles County"}})
	properties = append(properties, Property{ID: "7", Type: "Mansion", Bedrooms: 15, Price: 35000000, Area: &Area{Street: "Loma Vista Drive", City: "Los Angeles", Province: "Los Angeles County"}})
	properties = append(properties, Property{ID: "8", Type: "Mansion", Bedrooms: 10, Price: 25999999, Area: &Area{Street: "Loma Vista Drive", City: "Los Angeles", Province: "Los Angeles County"}})

	//Route Handlers / Endpoints
	r.HandleFunc("/api/properties", getProperties).Methods("GET")
	r.HandleFunc("/api/properties/{id}", getProperty).Methods("GET")
	r.HandleFunc("/api/properties", createProperty).Methods("POST")
	r.HandleFunc("/api/properties/{id}", updateProperty).Methods("PUT")
	r.HandleFunc("/api/properties/{id}", deleteProperty).Methods("DELETE")
	r.HandleFunc("/api/healthcheck", healthTestEndpoint)

	log.Fatal(http.ListenAndServe("localhost:9000", r))
}
