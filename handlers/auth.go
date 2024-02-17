package handlers

import (
	"embed"
	"errors"
	"html/template"
	"log"
	"net/http"
	"restdis/repositories"
	"restdis/types"
	"restdis/utils"
)

func RenderLoginPage(fs *embed.FS) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ, err := template.ParseFS(fs, "views/base.html", "views/login.html")
		if err != nil {
			log.Printf("Error while loading templates: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		templ.ExecuteTemplate(w, "base", nil)
		return
	})
}

func HandleLogin(users *repositories.Users) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		if username == "" {
			w.Write([]byte("<h1>Invalid credentials</h1>"))
			return
		}
		password := r.FormValue("password")
		if password == "" {
			w.Write([]byte("<h1>Invalid credentials</h1>"))
			return
		}
		dbUser, err := users.GetByUsername(username)
		if err != nil {
			if errors.Is(err, types.ErrUserNotFound) {
				w.Write([]byte("<h1>Invalid credentials</h1>"))
				return
			}
			log.Println(err.Error())
			w.Write([]byte("<h1>Something went wrong. Please retry</h1>"))
			return
		}
		if !utils.HashMatchesString(dbUser.Password, password) {
			w.Write([]byte("<h1>Invalid credentials</h1>"))
			return
		}
		w.Write([]byte("<h1>You in that shi!!!</h1>"))
		return
	})
}

func HandleChangePassword() http.Handler {
	return nil
}
