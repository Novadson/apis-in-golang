package movies

import "github.com/m4rw3r/uuid"

func GetAllMoviesBO() (moviesList []Movies, err error) {
	return GetAllMoviesDAO()
}

func GetMoviesByIdBO(Id uuid.UUID) (moviesList []Movies, err error) {
	return GetMoviesByIdDAO(Id)
}

func SaveOrUpdateBO(Movies Movies) (sucessMessage string, err error) {
	return SaveOrUpdateDAO(Movies)
}

func DeleteBO(Id uuid.UUID) (sucessMessage string, err error) {
	return DeleteDAO(Id)
}
