package main

import (
	"Halikha/database"
	"Halikha/types"
	"embed"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

//go:embed public
var FS embed.FS

func main() {
	if err := database.InitializeDatabase(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()

	// handle static files
	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	// // Serve Bootstrap CSS file
	// app.Static("/bootstrap/css", "./node_modules/bootstrap/dist/css")

	// // Serve Bootstrap JS files
	// app.Static("/bootstrap/js", "./node_modules/bootstrap/dist/js")
	// app.Static("/bootstrap/js/popper", "./node_modules/@popperjs/core/dist/umd")

	// // Serve other static files
	// app.Static("/", "./public")
	// app.Static("/", "./public/css")

	app.Get("/", login)
	app.Post("/register", register)
	app.Post("/authenticate", authenticate)
	app.Post("/upload", uploadImage)
	app.Get("/design", design)
	app.Post("/save-artwork", saveArtwork)

}

