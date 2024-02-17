package server

import (
	"embed"
	"net/http"
	"restdis/handlers"
	"restdis/repositories"
)

func RegisterRoutes(mux *http.ServeMux, fs *embed.FS, users *repositories.Users) {
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	mux.Handle("GET /login", handlers.RenderLoginPage(fs))
	mux.Handle("POST /auth/login", handlers.HandleLogin(users))
}
