package data

import (
	"accessory.nurtaymalika.com/internal/validator"
)

type Wish struct {
	ID    int64  `json:"id""`
	Title string `json:"title"`
	Price int64  `json:"price"`
}

func ValidateWish(v *validator.Validator, wish *Wish) {
	v.Check(wish.Title != "", "title", "must be provided")
	v.Check(len(wish.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(wish.Price < 0, "price", "must be more than 0")
	v.Check(wish.Price > 500000, "price", "must not be more than 500000")
}
