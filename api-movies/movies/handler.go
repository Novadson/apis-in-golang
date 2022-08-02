package movies

import (
	"fmt"
	"net/http"
)

func GetAllMoviesHandler(w http.ResponseWriter, r *http.Request) {

	//json.NewDecoder
}

func GetMoviesByIHandler(w http.ResponseWriter, r *http.Request) {
	/*Id := r.URL.Query().Get("id") *SECOND WAY TO GET PARAMETER VALUE* */
	Id := r.FormValue("id")
	fmt.Println(Id)
	//json.NewDecoder
}

func SaveOrUpdateHandler(w http.ResponseWriter, r *http.Request) {

	//json.NewDecoder
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	fmt.Println(Id)
	/*Id := r.FormValue("id") *SECOND WAY TO GET PARAMETER VALUE* */
	//json.NewDecoder
}
