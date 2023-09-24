package main

import (
	// "github.com/redis/go-redis/v9"
	"database/sql"
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Rights   string `json:"rights"`
}

type login_payload struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type Config struct {
	JwtSecret string `db:"jwt_secret"`
}

func main() {
	db, err := sqlx.Open("sqlite3", "./data.db")
	if err != nil {
		fmt.Printf("Error while connecting to local sqlite database %s", err.Error())
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error while connecting to local sqlite database %s", err.Error())
		os.Exit(1)
	}
	r := chi.NewRouter()

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
			var body = login_payload{}
			err := json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				println(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			var user = User{}
			err = db.Get(&user, `
        select * from users where username=:username;
      `, body.Username)
			if err != nil {
				if err == sql.ErrNoRows {
					http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
					return
				}
				println(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			match := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
			if match != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			var config = Config{}
			err = db.Get(&config, "select * from config")
			if err != nil {
				println(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			t, err := GenJWT(user.Username, config.JwtSecret)
			if err != nil {
				println(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			resp, err := json.Marshal(map[string]string{
				"token": t,
			})
			if err != nil {
				println(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			w.Write(resp)
		})
	})

	println("Restdis server launched!")
	http.ListenAndServe(":8080", r)
}
