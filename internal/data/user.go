package data

import (
	"accessory.nurtaymalika.com/internal/validator"
)

type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Username != "", "title", "must be provided")
	v.Check(len(user.Username) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(user.Password != "", "title", "must be provided")
	v.Check(len(user.Password) <= 500, "title", "must not be more than 500 bytes long")
}
