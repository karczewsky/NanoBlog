package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var database *sql.DB

func connectToDb() *sql.DB {
	connStr := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(fmt.Sprintf("Error during sql.Open to DB: %v\n", err))
	}

	return db
}

func initDB() {
	_, err := database.Query(`
		CREATE TABLE IF NOT EXISTS articles
		(id SERIAL PRIMARY KEY,
		title text NOT NULL,
		body text NOT NULL
		)
	`)

	if err != nil {
		panic(err)
	}
}
