package server

import (
	"embed"
	"net/http"
	"restdis/repositories"
)

func NewServer(viewsFS *embed.FS, users *repositories.Users) http.Handler {
	mux := http.NewServeMux()
	RegisterRoutes(mux, viewsFS, users)
	return mux
}
