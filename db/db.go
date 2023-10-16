package db

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectToDB() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", "/home/thewisepigeon/code/restdis/data.db")
	if err != nil {
		println("Error while connecting to database:", err.Error())
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		println("Error while connecting to database:", err.Error())
		os.Exit(1)
	}
	return db
}
