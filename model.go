package main

var superheros []SuperHero
var superhero SuperHero
var responseObject Response
var resource SuperHero

type SuperHero struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Biography  Biography  `json:"biography"`
	PowerStats PowerStats `json:"powerstats"`
	Work       Work       `json:"work"`
	Image      Image      `json:"image"`
	UUID       string     `json:"uuid"`
}
type Biography struct {
	Fullname string `json:"full-name"`
}
type Image struct {
	Url string `json:"url"`
}
type PowerStats struct {
	Power        string `json:"power"`
	Intelligence string `json:"intelligence"`
}
type Work struct {
	Occupation string `json:"occupation"`
}
type Response struct {
	Name       string     `json:"name"`
	PowerStats PowerStats `json:"powerstats"`
	Biography  Biography  `json:"biography"`
	Work       Work       `json:"work"`
	Image      Image      `json:"image"`
}
