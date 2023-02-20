package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	fmt.Println(params["id"])
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

var movies []Movie

func main() {
	movies = append(movies, Movie{
		ID: "1", Isbn: "69",
		Title: "Movie One",
		Director: &Director{
			FirstName: "alex",
			LastName:  "bob"}})

	movies = append(movies, Movie{
		ID: "2", Isbn: "420",
		Title: "Movie Two",
		Director: &Director{
			FirstName: "foo",
			LastName:  "bar"}})

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovies).Methods("GET")
	// r.HandleFunc("/movies", createMovies).Methods("POST")
	// r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", r)
}
