package data

import (
	"accessory.nurtaymalika.com/internal/validator"
	"database/sql"
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
type AccessoryModel struct {
	DB *sql.DB
}

// Add a placeholder method for inserting a new record in the movies table.
func (m AccessoryModel) Insert(accessory *Accessory) error {
	query := `
INSERT INTO accessory (title, year, price) VALUES (ipad, 2020, 4000)
RETURNING id`
	// Create an args slice containing the values for the // the movie struct. Declaring this slice immediately // make it nice and clear *what values are being used args := []any{movie.Title, movie.Year, movie.Runtime,
	// Use the QueryRow() method to execute the SQL query
	// passing in the args slice as a variadic parameter and scanning the system-
	// generated id, created_at and version values into the movie struct.
	return m.DB.QueryRow(query).Scan(&accessory.ID)
}

// Add a placeholder method for fetching a specific record from the movies table.
func (m AccessoryModel) Get(id int64) (*Accessory, error) {
	return nil, nil
}

// Add a placeholder method for updating a specific record in the movies table.
func (m AccessoryModel) Update(accessory *Accessory) error {
	return nil
}

// Add a placeholder method for deleting a specific record from the movies table.
func (m AccessoryModel) Delete(id int64) error {
	return nil
}

type MockAccessoryModel struct{}

func (m MockAccessoryModel) Insert(accessory *Accessory) error {
	// Mock the action...
	return nil
}
func (m MockAccessoryModel) Get(id int64) (*Accessory, error) {
	// Mock the action...
	return nil, nil
}
func (m MockAccessoryModel) Update(accessory *Accessory) error {
	// Mock the action...
	return nil
}
func (m MockAccessoryModel) Delete(id int64) error {
	// Mock the action...
	return nil
}
