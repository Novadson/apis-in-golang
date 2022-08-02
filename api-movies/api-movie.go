package main

import (
	"api-movies/movies"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/movies/", movies.GetAllMoviesHandler).Methods("GET")
	router.HandleFunc("/movies/{movieid}", movies.GetMoviesByIHandler).Methods("GET")
	router.HandleFunc("/movies/{id}", movies.GetMoviesByIHandler).Methods("DELETE")
	router.HandleFunc("/movies/", movies.SaveOrUpdateHandler).Methods("POST")

	fmt.Println("Server listening at port: 9000")
	log.Fatal(http.ListenAndServe(":9000", router))

}
