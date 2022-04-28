package main

import (
	// "crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"ISBN"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Name string `json:"name"`
}

var movies []Movie

func main() {
	// http.G
	route := mux.NewRouter()
	movies = append(movies, Movie{Id: "1", Isbn: "12345", Title: "New Book", Director: &Director{Name: "John Doe"}})

	route.HandleFunc("/", getMovies).Methods("GET")
	route.HandleFunc("/{id}", getMovies).Methods("GET")
	route.HandleFunc("/", createMovie).Methods("POST")
	route.HandleFunc("/{id}", updateMovie).Methods("PUT")
	route.HandleFunc("/{id}", deleteMovie).Methods("DELETE")

	http.Handle("/", route)

	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", route))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, movie := range movies {
		if movie.Id == params["id"] {
			json.NewEncoder(w).Encode(movies[i])
			break
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(1000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range movies {
		if item.Id == params["id"] {
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = item.Id
			movies[i] = movie
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, movie := range movies {
		if movie.Id == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)

}
