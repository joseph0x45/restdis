package server

import (
	"embed"
	"net/http"
)

func NewServer(viewsFS *embed.FS) http.Handler {
	mux := http.NewServeMux()
	RegisterRoutes(mux, viewsFS)
	return mux
}
