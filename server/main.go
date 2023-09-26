package main

import (
	// "github.com/redis/go-redis/v9"
	"database/sql"
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
)

type User struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Rights   string `json:"rights" db:"rights"`
}

type BlacklistedTokens struct {
	Token string `json:"token" db:"token"`
}

type login_payload struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type Config struct {
	JwtSecret string `db:"jwt_secret"`
}

var Db *sqlx.DB

func main() {
	Db, err := sqlx.Open("sqlite3", "./data.db")
	if err != nil {
		fmt.Printf("Error while connecting to local sqlite database %s", err.Error())
		os.Exit(1)
	}
	err = Db.Ping()
	if err != nil {
		fmt.Printf("Error while connecting to local sqlite database %s", err.Error())
		os.Exit(1)
	}
	r := chi.NewRouter()
	r.Use(middleware.DefaultLogger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
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
			err = Db.Get(&user, `
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
			err = Db.Get(&config, "select * from config")
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

	r.Route("/tokens", func(r chi.Router) {
		r.Use(Auth(""))
		r.Post("/", GenerateAccessToken)
		r.Delete("/", BlacklistToken)
	})

	r.Route("/", func(r chi.Router) {
		r.Use(Auth("w"))
	})

	println("Restdis server launched!")
	http.ListenAndServe(":8080", r)
}
