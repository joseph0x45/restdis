package server

import (
	"embed"
	"net/http"
	"restdis/handlers"
)

func RegisterRoutes(mux *http.ServeMux, fs *embed.FS) {
	mux.Handle("/login", handlers.RenderLoginPage(fs))
}
