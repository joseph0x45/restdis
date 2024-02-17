package server

import (
	"embed"
	"net/http"
	"restdis/handlers"
)

func RegisterRoutes(mux *http.ServeMux, fs *embed.FS) {
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	mux.Handle("GET /login", handlers.RenderLoginPage(fs))
}
