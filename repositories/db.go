package repositories

import (
	"database/sql"
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
	return db
}
