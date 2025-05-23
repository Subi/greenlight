package main

import (
	"fmt"
	"net/http"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdParams(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintln(w, "show movie details", id)
}
