package main

type SuperHero struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Fullname     string `json:"fullname"`
	Intelligence int    `json:"intelligence"`
	Power        int    `json:"power"`
	Occupation   string `json:"occupation"`
	Image        string `json:"image"`
	Work         string `json:"work"`
	UUID         string `json:"uuid"`
}

var superheros []SuperHero

type Response struct {
	Name         string `json:"name"`
	Fullname     string `json:"fullname"`
	Intelligence int    `json:"intelligence"`
	Power        int    `json:"power"`
	Occupation   string `json:"occupation"`
	Image        string `json:"image"`
	Work         string `json:"work"`
}
