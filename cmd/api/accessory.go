package main

import (
	"accessory.nurtaymalika.com/internal/data"
	_ "encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Add a createMovieHandler for the "POST /v1/movies" endpoint. For now we simply
// return a plain-text placeholder response.
func (app *application) createAccessoryHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Title   string `json:"title"`
		Color   string `json:"color"`
		Country string `json:"country"`
		Year    string `json:"year"`
		Device  string `json:"device"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showAccessoryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Create a new instance of the Movie struct, containing the ID we extracted from
	//the URL and some dummy data. Also notice that we deliberately haven't set a
	// value for the Year field.
	accessory := data.Accessory{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Case",
		Color:     "Black",
		Country:   "China",
		Year:      2020,
		Device:    "Iphone 13",
	}
	// Encode the struct to JSON and send it as the HTTP response.
	err = app.writeJSON(w, http.StatusOK, envelope{"accessory": accessory}, nil)
	if err != nil {
		// Use the new serverErrorResponse() helper.
		app.serverErrorResponse(w, r, err)
	}
}
