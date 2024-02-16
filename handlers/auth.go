package handlers

import (
	"html/template"
	"net/http"
)

func RenderLoginPage(tpl *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from the login page"))
		return
	})
}
