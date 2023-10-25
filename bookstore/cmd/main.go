package main

import (
	"github.com/gorilla/mux"
	"github.com/tudosm/go-projects/bookstore/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
