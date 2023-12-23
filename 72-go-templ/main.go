package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

//go:embed public
var staticFiles embed.FS

func main() {
	staticFS := fs.FS(staticFiles)
	staticContent, err := fs.Sub(staticFS, "public")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.FS(staticContent))))
	http.Handle("/", templ.Handler(HomePage("Vasya")))

	addr := "localhost:8080"
	log.Printf("listening on http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
