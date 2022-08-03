package dataservices

import (
	app "api-movies/app-message"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	HOST      = "postgres"
	USER_NAME = "postgres"
	DB_NAME   = "postgres"
	PASSWORD  = "postgres"
	PORT      = 5432
)

func SetUpDb() (db *sqlx.DB) {
	dbInfo := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d",
		HOST, USER_NAME, DB_NAME, PASSWORD, PORT)

	db, err := sqlx.Open("postgres", dbInfo)
	app.CheckErroMessage(err)

	err = db.Ping()
	app.CheckErroMessage(err)

	defer db.Close()

	return
}
