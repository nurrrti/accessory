package main

import (
	"accessory.nurtaymalika.com/internal/data"
	"accessory.nurtaymalika.com/internal/validator"
	"fmt"
	"net/http"
)

func (app *application) createWishHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
		Price int64  `json:"price"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	wish := &data.Wish{
		Title: input.Title,
		Price: input.Price,
	}
	v := validator.New()
	if data.ValidateWish(v, wish); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}
func (app *application) showWishHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	wish := data.Wish{
		ID:    id,
		Title: "Case",
		Price: 2000,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"wish": wish}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
