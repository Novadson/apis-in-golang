package dataservices

import (
	app "api-movies/app-message"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	HOST      = "postgres"
	USER_NAME = "postgres"
	DB_NAME   = "postgres"
	PASSWORD  = "postgres"
	PORT      = 5432
)

func SetUpDb() *sql.DB {
	dbInfo := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d",
		HOST, USER_NAME, DB_NAME, PASSWORD, PORT)

	db, err := sql.Open("postgres", dbInfo)
	app.CheckErroMessage(err)

	defer db.Close()

	return db
}
