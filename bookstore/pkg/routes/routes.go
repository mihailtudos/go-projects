package routes

import (
	"github.com/gorilla/mux"
	"github.com/tudosm/go-projects/bookstore/pkg/controllers"
	"net/http"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/books", controllers.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{bookId}", controllers.ShowBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{bookId}", controllers.UpdateBooks).Methods(http.MethodPut)
	router.HandleFunc("/books/{bookId}", controllers.DeleteBooks).Methods(http.MethodDelete)

}
