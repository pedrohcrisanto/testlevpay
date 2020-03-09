package main

type SuperHero struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var superheros []SuperHero

type Response struct {
	Name string `json:"name"`
}
