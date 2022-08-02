package movies

import (
	"time"

	"github.com/m4rw3r/uuid"
)

type Movies struct {
	Id         uuid.UUID  `json:"id" db:"id"`
	MovieName  string     `json:"movie_name" db:"movie_name"`
	Year       time.Time  `json:"year" db:"year"`
	Duration   time.Time  `json:"duration" db:"duration"`
	CategoryId uuid.UUID  `json:"category_id" db:"category_id"`
	DirectorId uuid.UUID  `json:"director_id" db:"director_id"`
	StudioId   uuid.UUID  `json:"studio_id" db:"studio_id"`
	Categories Categories `json:"Categories"`
}

type Categories struct {
	Id           uuid.UUID `json:"id" db:"id"`
	CategoryName string    `json:"category_name" db:"category_name"`
}

type Studios struct {
	Id         uuid.UUID `json:"id" db:"id"`
	StudioName string    `json:"studio_name" db:"studio_name"`
	Address    string    `json:"address" db:"address"`
}

type Directors struct {
	Id           uuid.UUID `json:"id" db:"id"`
	DirectorName string    `json:"director_name" db:"director_name"`
	Address      string    `json:"address" db:"address"`
	CategoryId   string    `json:"category_id"`
	Categories   string    `json:"Categories"`
}

type Actors struct {
	Id          uuid.UUID `json:"id" db:"id"`
	Address     string    `json:"address" db:"address"`
	Sex         string    `json:"sex" db:"sex"`
	DateOfBirth string    `json:"date_of_birth" db:"date_of_birth"`
}
