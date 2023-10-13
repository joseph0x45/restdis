package config

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type Config struct {
	DbLocation string
	JWTSecret  string
}

func GetConfig() {

}

func Init(db_location string) {
	db, err := sqlx.Open("sqlite3", db_location)
	if err != nil {
		println("Something went wrong : ", err.Error())
		os.Exit(1)
	}
	init_query := `
    create table if not exists users(
      id integer primary key autoincrement,
      username text not null unique,
      password text not null
    );

    create table if not exists configs(
      id integer primary key autoincrement,
      env text default 'main' unique,
      db_location text not null,
      jwt_secret text not null
    );
  `
	_, err = db.Exec(init_query)
	if err != nil {
		println("Something went wrong : ", err.Error())
		os.Exit(1)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte("root"), 4)
	if err != nil {
		println("Something went wrong : ", err.Error())
		os.Exit(1)
	}
	_, err = db.Exec(
		`insert into users(username, password) values($1, $2)`,
		"root",
		string(hash),
	)
	if err != nil {
		println("Something went wrong : ", err.Error())
		os.Exit(1)
	}
	_, err = db.Exec(
		`insert into configs(db_location, jwt_secret) values($1, $2)`,
		db_location,
		"randomJWTsecret",
	)
	if err != nil {
		println("Something went wrong : ", err.Error())
		os.Exit(1)
	}
}
