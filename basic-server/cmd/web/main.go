package main

import (
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)

	log.Printf("starting server on %s", PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}
