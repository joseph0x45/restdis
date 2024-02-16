package server

import (
	"html/template"
	"net/http"
)

func NewServer(tmpl template.Template) http.Handler {
	mux := http.NewServeMux()
	RegisterRoutes(mux, tmpl)
	return mux
}
