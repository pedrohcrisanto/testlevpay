package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func getSuperHeros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Query("SELECT id, name, fullname, intelligence, power, occupation, image, uuid from superheros")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&superhero.ID, &superhero.Name, &superhero.Biography.Fullname,
			&superhero.PowerStats.Intelligence, &superhero.PowerStats.Power, &superhero.Work.Occupation,
			&superhero.Image.Url, &superhero.UUID)
		if err != nil {
			panic(err.Error())
		}
		switch err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
			return
		case nil:
			json.NewEncoder(w).Encode(superhero)
		default:
			panic(err)
		}
	}
}

func createSuperHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	variable := params["id"]
	url := "https://superheroapi.com/api/" + superheroapikey + "/" + variable

	response, err := http.Get(url)
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
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
	INSERT INTO superheros (name, fullname, intelligence, power, occupation, image, uuid)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, responseObject.Name, responseObject.Biography.Fullname,
		responseObject.PowerStats.Intelligence, responseObject.PowerStats.Power,
		responseObject.Work.Occupation, responseObject.Image.Url, uuid.New()).Scan(&id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode("Criado com Sucesso!")
}

func getSuperHero(w http.ResponseWriter, r *http.Request) {
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

	sqlStatement := `SELECT * FROM superheros WHERE id=$1;`
	row := db.QueryRow(sqlStatement, params["id"])
	err = row.Scan(&superhero.ID, &superhero.Name, &superhero.Biography.Fullname, &superhero.PowerStats.Intelligence,
		&superhero.PowerStats.Power, &superhero.Work.Occupation, &superhero.Image.Url, &superhero.UUID)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		json.NewEncoder(w).Encode(superhero)
	default:
		panic(err)
	}

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
	var resource SuperHero
	d := json.NewDecoder(r.Body)
	d.Decode(&resource)
	defer r.Body.Close()
	sqlStatement := `
	UPDATE superheros
	SET name = $2
	WHERE id = $1;`
	_, err = db.Exec(sqlStatement, params["id"], resource.Name)
	json.NewEncoder(w).Encode("Nome Atualizado com Sucesso!")
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
	json.NewEncoder(w).Encode("Deletado com Sucesso!")
}

func getSearch(w http.ResponseWriter, r *http.Request) {
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

	sqlStatement := `SELECT * FROM superheros WHERE LOWER(name)=$1 OR uuid=$1;`
	var superhero SuperHero
	row := db.QueryRow(sqlStatement, strings.ToLower(params["query"]))
	err = row.Scan(&superhero.ID, &superhero.Name, &superhero.Biography.Fullname, &superhero.PowerStats.Intelligence,
		&superhero.PowerStats.Power, &superhero.Work.Occupation, &superhero.Image.Url, &superhero.UUID)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		json.NewEncoder(w).Encode(superhero)
	default:
		panic(err)
	}
}
