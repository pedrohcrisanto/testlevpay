package main

type SuperHero struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Biography   Biography  `json:"biography"`
	PowerStatus PowerStats `json:"powerstats"`
	Work        Work       `json:"work"`
	Image       Image      `json:"image"`
	UUID        string     `json:"uuid"`
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

var superheros []SuperHero

type Response struct {
	Name        string     `json:"name"`
	PowerStatus PowerStats `json:"powerstats"`
	Biography   Biography  `json:"biography"`
	Work        Work       `json:"work"`
	Image       Image      `json:"image"`
}
