package data

import (
	"accessory.nurtaymalika.com/internal/validator"
	"time"
)

type Accessory struct {
	ID        int64     `json:"id""`
	CreatedAt time.Time `json:"-""`
	Title     string    `json:"title"`
	Color     string    `json:"color"`
	Country   string    `json:"country"`
	Device    string    `json:"device"`
	Year      int32     `json:"year"`
	Price     int64     `json:"price"`
}

// Add a createMovieHandler for the "POST /v1/movies" endpoint. For now we simply
// return a plain-text placeholder response.
func ValidateAccessory(v *validator.Validator, accessory *Accessory) {
	v.Check(accessory.Title != "", "title", "must be provided")
	v.Check(len(accessory.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(accessory.Year != 0, "year", "must be provided")
	v.Check(accessory.Year >= 2000, "year", "must be greater than 2000")
	v.Check(accessory.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(accessory.Color != "", "Color", "must be provided")

	v.Check(accessory.Device != "", "Device", "must be provided")

	v.Check(accessory.Price != 0, "price", "must be provided")
	v.Check(accessory.Price >= 100, "price", "must be greater than 100")
	v.Check(accessory.Price <= 50000, "price", "must be lower than 50000")
}

// Unique integer ID for the movie
// Timestamp for when the movie is added to our database
// Movie title
// Movie release year
// Movie runtime (in minutes)
// Slice of genres for the movie (romance, comedy, etc.)
// The version number starts at 1 and will be incremented each // time the movie information is updated
