package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/at-tev/moon/pkg/models/postgresql"

	"github.com/getsentry/sentry-go"
	_ "github.com/lib/pq"
)

type application struct {
	reviews       *postgresql.ReviewModel
	templateCache map[string]*template.Template
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	addr := flag.String("addr", ":8080", "Сетевой адресс HTTP")
	dsn := flag.String("dsn", "user=postgres password=Mint_22 dbname=moon sslmode=disable", "Строка подключения к Postgresql")
	flag.Parse()

	log.Println("Sentry init")
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://7b66501c8d42441f9e7b6520c0f07923@o1387515.ingest.sentry.io/4504577483210752",
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Printf("sentry.Init: %s", err)
	}

	log.Println("Connection to DB")
	db, err := openDB(*dsn)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	log.Println("Creating template cashe")
	templateCache, err := newTemplateCache("../../ui/html")
	if err != nil {
		log.Println(err)
	}

	app := &application{
		reviews:       &postgresql.ReviewModel{DB: db},
		templateCache: templateCache,
	}

	mux := app.routes()
	srv := &http.Server{
		Addr:    *addr,
		Handler: mux,
	}

	log.Printf("Server is lunched on address %s", *addr)
	log.Fatal(srv.ListenAndServe())
}
