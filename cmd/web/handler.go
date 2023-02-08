package main

import (
	"net/http"
	"strings"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.clientError(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		app.clientError(w, http.StatusForbidden)
		return
	}

	lang := strings.Split(strings.Split(r.Header.Get("Accept-Language"), ";")[0], ",")[0]

	tmplData, err := app.reviews.Random4(lang)

	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("Content-Security-Policy",
		"default-src 'self' captcha-api.yandex.ru")
	app.render(w, r, "home.page.tmpl", &tmplData)
}

func (app *application) allReviews(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/reviews" {
		app.clientError(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		app.clientError(w, http.StatusForbidden)
		return
	}

	lang := r.Header.Get("Accept-Language")

	tmplData, err := app.reviews.AllReviews(lang[0:2])
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	app.render(w, r, "reviews.page.tmpl", &tmplData)
}

func (app *application) createReview(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/feedback" {
		app.clientError(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodPost:
		err := app.reviews.Insert(r.FormValue("user"), r.FormValue("review"))
		if err != nil {
			app.serverError(w, err)
			return
		}

		http.Redirect(w, r, "/#reviews", http.StatusSeeOther)

	case http.MethodGet:
		w.Header().Set("Content-Type", "text/html")
		app.render(w, r, "create_review.page.tmpl", nil)

	default:
		app.clientError(w, http.StatusForbidden)
		return
	}
}
