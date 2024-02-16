package main

import (
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net"
	"net/http"
	"restdis/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	tmpl := template.Template{}
	srv := server.NewServer(tmpl)
	server := &http.Server{
		Addr:    net.JoinHostPort("localhost", "8080"),
		Handler: srv,
	}
	log.Println("Server launched on port 8080")
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
