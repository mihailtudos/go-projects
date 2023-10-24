package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	mihail := Director{"Mihail", "Tudos"}

	movies = append(movies,
		Movie{uuid.New().String(), strings.Replace(uuid.New().String(), "-", "", -1), "My best friend", &mihail},
		Movie{uuid.New().String(), strings.Replace(uuid.New().String(), "-", "", -1), "Last memory", &mihail},
	)

	r.HandleFunc("/movies", getMovies).Methods(http.MethodGet)
	r.HandleFunc("/movies/{id}", getMovie).Methods(http.MethodGet)
	r.HandleFunc("/movies", createMovie).Methods(http.MethodPost)
	r.HandleFunc("/movies/{id}", updateMovie).Methods(http.MethodPut)
	r.HandleFunc("/movies/{id}", deleteMovie).Methods(http.MethodDelete)

	fmt.Println("Starting server at port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			var m Movie
			_ = json.NewDecoder(r.Body).Decode(&m)
			m.ID = params["id"]
			movies = append(movies[:index], movies[index+1:]...)
			movies = append(movies, m)
			_ = json.NewEncoder(w).Encode(m)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = uuid.New().String()
	movies = append(movies, movie)
	_ = json.NewEncoder(w).Encode(movie)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var movie *Movie

	for index, item := range movies {
		if item.ID == params["id"] {
			movie = &movies[index]
			break
		}
	}

	if movie == nil {
		http.Error(w, "Movie not found.", http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(movie)
	if err != nil {
		http.Error(w, "Something went wrong, try again later.", http.StatusInternalServerError)
	}
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			return
		}
	}

	json.NewEncoder(w).Encode(movies)
}
