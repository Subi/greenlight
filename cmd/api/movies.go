package main

import (
	"fmt"
	"github.com/subi/greenlight/internal/data"
	"net/http"
	"time"
)

type input struct {
	Title   string   `json:"title"`
	Year    int      `json:"year"`
	Runtime int32    `json:"runtime"`
	Genres  []string `json:"genres"`
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	input := &input{}

	err := app.readJSON(r, w, input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParams(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Focus",
		Year:      2015,
		Runtime:   120,
		Genres:    []string{"drama", "crime", "mystery"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
