package db

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectToDB() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", "config.DbLocation")
	if err != nil {
		println("Error while connecting to database:", err.Error())
		os.Exit(1)
	}
	return db
}
