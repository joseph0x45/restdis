package handlers

import (
	"embed"
	"html/template"
	"log"
	"net/http"
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
