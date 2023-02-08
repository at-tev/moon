package main

import (
	"log"
	"net/http"
	"path/filepath"
)

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		log.Println(err)
	}

	if s.IsDir() {
		index := filepath.Join(path, "index.html")

		if _, err = nfs.fs.Open(index); err != nil {
			closeErr := f.Close()

			if closeErr != nil {
				return nil, err
			}

			return nil, err
		}
	}

	return f, nil
}

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/reviews", app.allReviews)
	mux.HandleFunc("/feedback", app.createReview)

	staticHandler := http.FileServer(neuteredFileSystem{http.Dir("../../ui")})
	mux.Handle("/static/", staticHandler)

	return mux
}
