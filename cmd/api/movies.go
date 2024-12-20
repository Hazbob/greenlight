package main

import (
	"fmt"
	"github.com/hazbob/greenlight/internal/data"
	"net/http"
	"time"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "new movie creater")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Time{},
		Title:     "Pulp Fiction",
		Year:      0,
		Runtime:   "90",
		Genres:    []string{"Comedy", "Action"},
		Version:   1,
	}

	err = app.writeJson(w, http.StatusOK, movie, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not proces your request", http.StatusInternalServerError)
		return
	}
}
