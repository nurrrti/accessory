package data

import (
	"errors"
)

// Define a custom ErrRecordNotFound error. We'll return this from our Get() method when // looking up a movie that doesn't exist in our database.
var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	// Set the Movies field to be an interface containing the methods that both the // 'real' model and mock model need to support.
	Accessory interface {
		Insert(accessory *Accessory) error
		Get(id int64) (*Accessory, error)
		Update(accessory *Accessory) error
		Delete(id int64) error
	}
}

// Create a helper function which returns a Models instance containing the mock models // only.
func NewMockModels() Models {
	return Models{
		Accessory: MockAccessoryModel{},
	}
}
