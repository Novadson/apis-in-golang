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

	router.HandleFunc("/movies/", movies.GetHandler).Methods("GET")

	fmt.Println("Server listening at port: 9000")
	log.Fatal(http.ListenAndServe(":9000", router))

}
