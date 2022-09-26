package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// stuct - key value pair

type Movie struct {
	ID       string    `json: "id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json : "firstname"`
	Lastname  string `json: "lastname"`
}

var movies []Movie

// get all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// get movie by ID
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// delete movie by ID
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)

	for index, item := range movies {
		if item.ID == parms["id"] {
			movies = append(movies[:index], movies[index+1:]...) // vImpt
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

// create movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// update movie
func updateMovie(w http.ResponseWriter, r *http.Request) {
	// set json content type
	w.Header().Set("Content-Type", "application/json")
	// params
	parms := mux.Vars(r)

	for index, item := range movies {
		if parms["id"] == item.ID {

			movies = append(movies[:index], movies[index+1:]...) // v - impt
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = parms["id"]
			movies = append(movies, movie)

			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "42337", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})

	movies = append(movies, Movie{ID: "2", Isbn: "42333", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/delete/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting Server at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
