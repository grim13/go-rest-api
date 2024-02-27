package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 5433
	user     = "go-rest-api"
	password = "qwerty123"
	dbname   = "go-rest-api"
)

func InitDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	DB, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	defer DB.Close()
	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime TIMESTAMP NOT NULL,
		user_id INTEGER,
		created_at TIMESTAMP not null,
		updated_at TIMESTAMP not null
	)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic(err)
	}
}
