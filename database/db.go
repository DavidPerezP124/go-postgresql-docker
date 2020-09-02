package database

import (
	"database/sql"
	"fmt"
	"log"

	"example.com/m/config"
)

//StartDB initializes the database
func StartDB() (db *sql.DB, err error) {
	m := config.Config()
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", m["DB_USER"], m["DB_PASSWORD"], m["DB_NAME"], m["PORT"])

	log.Printf("Postgres started at %s PORT", m["PORT"])

	return sql.Open("postgres", dbinfo)
}
