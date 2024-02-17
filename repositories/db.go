package repositories

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "restdis.db")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
  log.Println("Connected to SQLite database")
	return db
}
