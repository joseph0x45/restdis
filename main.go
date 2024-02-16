package main

import (
	"embed"
	"github.com/joho/godotenv"
	"log"
	"net"
	"net/http"
	"restdis/server"
)

//go:embed views
var views embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	srv := server.NewServer(&views)
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
