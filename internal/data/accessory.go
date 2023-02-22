package data

import (
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
}

// Unique integer ID for the movie
// Timestamp for when the movie is added to our database
// Movie title
// Movie release year
// Movie runtime (in minutes)
// Slice of genres for the movie (romance, comedy, etc.)
// The version number starts at 1 and will be incremented each // time the movie information is updated
