package main

import (
	"fmt"
	"net/http"

	"github.com/at-tev/moon/pkg/models"
	"github.com/getsentry/sentry-go"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	// trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	// app.errorLog.Output(2, trace)
	sentry.CaptureException(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, tmplData *[]*models.Review) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("there is no template %s", name))
		return
	}

	err := ts.Execute(w, tmplData)
	if err != nil {
		app.serverError(w, err)
	}
}
