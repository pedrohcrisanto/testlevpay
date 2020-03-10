package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func getSuperHeros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func createSuperHero(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	variable := params["id"]
	url := "https://superheroapi.com/api/10157313301043512/" + variable

	response, err := http.Get(url)
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
	INSERT INTO superheros (name, fullname, intelligence, power, occupation, image, work, uuid)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id`
	id := 0
	fmt.Println(responseObject)
	err = db.QueryRow(sqlStatement, responseObject.Name, "Nome Completo", 25, 40, "Trabalhador", "URL IMAGE", "Vendendor", uuid.New()).Scan(&id)
	if err != nil {
		panic(err)
	}
}

func getSuperHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func updateSuperHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
	UPDATE superheros
	SET name = $2
	WHERE id = $1;`
	_, err = db.Exec(sqlStatement, params["id"], "Pedro")
	if err != nil {
		panic(err)
	}

}

func deleteSuperHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
	DELETE FROM superheros
	WHERE id = $1;`
	_, err = db.Exec(sqlStatement, params["id"])
	if err != nil {
		panic(err)
	}

}
