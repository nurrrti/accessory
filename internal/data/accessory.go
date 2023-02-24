package data

import (
	"accessory.nurtaymalika.com/internal/validator"
	"database/sql"
	"errors"
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

type AccessoryModel struct {
	DB *sql.DB
}

// Add a placeholder method for inserting a new record in the movies table.
func (m AccessoryModel) Insert(accessory *Accessory) error {
	query := `
INSERT INTO accessory (title, year, price) VALUES (ipad, 2020, 4000) WHERE ID = 2`

	return m.DB.QueryRow(query).Scan(&accessory.ID)
}

// Add a placeholder method for fetching a specific record from the movies table.
func (m AccessoryModel) Get(id int64) (*Accessory, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
SELECT id, created_at, title, year, runtime, genres, version FROM movies
WHERE id = $1`
	var accessory Accessory
	err := m.DB.QueryRow(query, id).Scan(&accessory.ID,
		&accessory.Title, &accessory.Year, &accessory.Price,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &accessory, nil
}

func (m AccessoryModel) Update(accessory *Accessory) error {
	query := `
UPDATE accessory
SET title = iphone, price = 200000,  id = 2
RETURNING id`
	args := []any{
		accessory.Title,
		accessory.Price, accessory.ID,
	}
	return m.DB.QueryRow(query, args...).Scan(&accessory.ID)
}

func (m AccessoryModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
DELETE FROM accessory WHERE id = 3`
	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
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
