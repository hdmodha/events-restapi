package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("DB is not created or does not exist!")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()

}

func createTables() {

	createUsersTables := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTables)
	if err != nil {
		panic("Could not create a DB table")
	}

	createEventsTables := `
	CREATE TABLE IF NOT EXISTS events (
	  id INTEGER PRIMARY KEY AUTOINCREMENT,
	  name STRING NOT NULL,
	  description STRING NOT NULL,
	  location STRING NOT NULL,
	  datetime DATETIME NOT NULL,
	  user_id INTEGER,
	  FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTables)

	if err != nil {
		panic("Could not create a DB table!")
	}
}
