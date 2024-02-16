package server

import (
	"embed"
	"net/http"
)

func NewServer(fs *embed.FS) http.Handler {
	mux := http.NewServeMux()
	RegisterRoutes(mux, fs)
	return mux
}
