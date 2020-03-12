package main

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1992"
	dbname   = "testlevpay3"
)

func main() {
	router := mux.NewRouter()
	// superheros = append(superheros, SuperHero{ID: "1", Name: "Spider-Man"})
	router.HandleFunc("/superheros", getSuperHeros).Methods("GET")
	router.HandleFunc("/search/{query}", getSearch).Methods("GET")
	router.HandleFunc("/superherosapi/{id}", createSuperHero).Methods("POST")
	router.HandleFunc("/superheros/{id}", getSuperHero).Methods("GET")
	router.HandleFunc("/superheros/{id}", updateSuperHero).Methods("PUT")
	router.HandleFunc("/superheros/{id}", deleteSuperHero).Methods("DELETE")
	http.ListenAndServe(":8000", router)

}
