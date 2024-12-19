package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "new movie creater")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	_, err = fmt.Fprintf(w, "show the details of movie %d\n", id)
	if err != nil {
		fmt.Fprintf(w, "failure to write response")
	}
}
