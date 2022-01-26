package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Fruits struct {
	Name   string `json:"name" xml:"name"`
	Season string `json:"season" xml:"season"`
	Price  int32  `json:"price" xml:"price"`
}

type Games struct {
	Name string `name`
	Type string `type`
	Year int32  `year`
}

type Clothes struct {
	Type  string `type`
	Brand string `brand`
	Price int    `price`
}

func getAllFruits(w http.ResponseWriter, r *http.Request) {
	fruits := []Fruits{
		{"mango", "summer", 80},
		{"guava", "monsoon", 50},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(fruits)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fruits)
	}
}

func getAllGames(w http.ResponseWriter, r *http.Request) {
	games := []Games{
		{"Call Of duty", "multiplayer", 2008},
		{"PUBG", "FPS", 2019},
		{"Mobile Legends", "Moba", 2016},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func getAllClothes(w http.ResponseWriter, r *http.Request) {
	clothes := []Clothes{
		{"Pant", "CatsEye", 1800},
		{"Shirt", "Yellow", 1600},
		{"T-shirt", "Tanjim", 2200},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clothes)
}
