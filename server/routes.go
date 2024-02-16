package server

import (
	"html/template"
	"net/http"
	"restdis/handlers"
)

func RegisterRoutes(mux *http.ServeMux, tmpl template.Template) {
	mux.Handle("/login", handlers.RenderLoginPage(&tmpl))
}
