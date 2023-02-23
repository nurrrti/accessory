package main

import (
	"accessory.nurtaymalika.com/internal/data"
	"accessory.nurtaymalika.com/internal/validator"
	_ "encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createAccessoryHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string `json:"title"`
		Color   string `json:"color"`
		Country string `json:"country"`
		Year    int32  `json:"year"`
		Device  string `json:"device"`
		Price   int64  `json:"price"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	// Copy the values from the input struct to a new Movie struct.
	accessory := &data.Accessory{
		Title:   input.Title,
		Color:   input.Color,
		Country: input.Country,
		Year:    input.Year,
		Device:  input.Device,
		Price:   input.Price,
	}
	// Initialize a new Validator.
	v := validator.New()
	// Call the ValidateMovie() function and return a response containing the errors if // any of the checks fail.
	if data.ValidateAccessory(v, accessory); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

// Use the Check() method to execute our validation checks. This will add the
// provided key and error message to the errors map if the check does not evaluate // to true. For example, in the first line here we "check that the title is not
// equal to the empty string". In the second, we "check that the length of the title // is less than or equal to 500 bytes" and so on.

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
		Price:     2000,
	}
	// Encode the struct to JSON and send it as the HTTP response.
	err = app.writeJSON(w, http.StatusOK, envelope{"accessory": accessory}, nil)
	if err != nil {
		// Use the new serverErrorResponse() helper.
		app.serverErrorResponse(w, r, err)
	}
}
