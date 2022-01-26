package main

import (
	"log"
	"net/http"
)

func start() {
	http.HandleFunc("/fruits", getAllFruits)
	http.HandleFunc("/games", getAllGames)
	http.HandleFunc("/clothes", getAllClothes)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
