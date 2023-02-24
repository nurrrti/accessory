package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Accessory interface {
		Insert(accessory *Accessory) error
		Get(id int64) (*Accessory, error)
		Update(accessory *Accessory) error
		Delete(id int64) error
	}
}

func NewAccessoryModels(db *sql.DB) Models {
	return Models{
		Accessory: MockAccessoryModel{},
	}
}
