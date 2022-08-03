package movies

import (
	"github.com/m4rw3r/uuid"
	"github.com/mitchellh/mapstructure"
)

func GetAllMoviesDAO() (moviesList []Movies, err error) {
	return
}

func GetMoviesByIdDAO(Id uuid.UUID) (moviesList []Movies, err error) {
	return
}

func SaveOrUpdateMoviesDAO(Movies Movies) (sucessMessage string, err error) {

	Category := GetDependendyDAO(Movies.CategoryId)
	mapstructure.Decode(Category, Movies.Categories)
	return
}

func DeleteMoviesDAO(Id uuid.UUID) (sucessMessage string, err error) {
	return
}

func SaveOrUpdateCategoriesDAO(Categories Categories) (succeMessage string, err error) {
	return
}

func SaveOrUpdateStudiosDAO(Studios Studios) (succeMessage string, err error) {
	return
}

func SaveOrUpdateDirectorsDAO(Directors Directors) (succeMessage string, err error) {
	return
}

func SaveOrUpdateActorsDAO(Actors Actors) (succeMessage string, err error) {
	return
}

func SaveOrUpdateDependencyDAO(Obj interface{}) (sucessMessage string, err error) {
	return
}

func GetDependendyDAO(Id uuid.UUID) interface {
}
