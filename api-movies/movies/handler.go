package movies

import (
	errMessa "api-movies/app-message"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAllMoviesHandler(w http.ResponseWriter, r *http.Request) {

	//json.NewDecoder
}

func GetMoviesByIHandler(w http.ResponseWriter, r *http.Request) {
	/*
		*router.HandleFunc("/movies"
		Id := r.URL.Query().Get("id")
		/*
		/*
		*router.HandleFunc("/movies/{id}"
		*params := mux.Vars(r) paras := params["movieid"]
		*THIRD WAY TO GET PARAMETER VALUE
	*/

	Id := r.FormValue("id")
	fmt.Println(Id)
	//json.NewDecoder
}

func SaveOrUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var Movies Movies

	err := json.NewDecoder(r.Body).Decode(&Movies)
	errMessa.CheckErroMessage(err)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	fmt.Println(Id)
	/*
		router.HandleFunc("/movies"
		Id := r.FormValue("id") *SECOND WAY TO GET PARAMETER VALUE* */

	/*
		router.HandleFunc("/movies/{id}"
		params := mux.Vars(r)
		 id := params["id"]
		*THIRD WAY TO GET PARAMETER VALUE* */

	//json.NewDecoder
}
