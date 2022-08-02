package movies

import "github.com/m4rw3r/uuid"

func GetAllMoviesDAO() (moviesList []Movies, err error) {
	return
}

func GetMoviesByIdDAO(Id uuid.UUID) (moviesList []Movies, err error) {
	return
}

func SaveOrUpdateDAO(Movies Movies) (sucessMessage string, err error) {
	return
}

func DeleteDAO(Id uuid.UUID) (sucessMessage string, err error) {
	return
}
