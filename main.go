package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed static/*
var embededFiles embed.FS

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(embededFiles, "static")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static/", http.FileServer(getFileSystem())).ServeHTTP(w, r)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := layout(index())
		component.Render(r.Context(), w)
	})

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", mux)
}
