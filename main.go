package main

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host            = "localhost"
	port            = 5432
	user            = "postgres"
	password        = "1992"
	dbname          = "testlevpay5"
	superheroapikey = "10157313301043512"
)

func main() {
	router := mux.NewRouter()
	// Listar todos os superheros
	router.HandleFunc("/superheros", getSuperHeros).Methods("GET")
	// Listar superhero a partir de seu name ou uuid
	router.HandleFunc("/search/{query}", getSearch).Methods("GET")
	// Cria um superhero a partir da API(Superheroapi.com) passando o ID e grava no Banco de Dados.
	router.HandleFunc("/superherosapi/{id}", createSuperHero).Methods("POST")
	// Lista o superhero do Banco de Dados a partir do seu ID
	router.HandleFunc("/superheros/{id}", getSuperHero).Methods("GET")
	// Realiza o Update do superhero no Banco de dados passando o ID como params e
	// o seu name no Body da request
	router.HandleFunc("/superheros/{id}", updateSuperHero).Methods("PUT")
	// Deleta o superhero do Banco de Dados passando o ID como params
	router.HandleFunc("/superheros/{id}", deleteSuperHero).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}
