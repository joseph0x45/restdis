package main

import (
	"embed"
	"log"
	"net"
	"net/http"
	"restdis/repositories"
	"restdis/server"

	"github.com/joho/godotenv"
)

//go:embed views
var views embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
  db := repositories.NewSQLiteConnection()
  usersRepop := repositories.NewUsersRepo(db)
	srv := server.NewServer(&views, usersRepop)
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
