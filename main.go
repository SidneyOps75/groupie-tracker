package main

import (
	"log"
	"net/http"

	"groupie-trackers/handlers"
)

// Fetch data from a URL and unmarshal it into a target interface

func main() {
	http.HandleFunc("GET /", handlers.HandleHomepage)
	http.HandleFunc("GET /artist/", handlers.HandleArtistDetail)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("GET /static/", http.StripPrefix("/static/", fs))

	fs2 := http.FileServer(http.Dir("images"))
	http.Handle("GET /images/", http.StripPrefix("/images/", fs2))

	log.Println("Server started at http://localhost:9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
