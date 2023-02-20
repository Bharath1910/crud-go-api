package main

import (
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
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/moviles/{id}", deleteMovie).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}
