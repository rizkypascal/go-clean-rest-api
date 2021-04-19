package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	ID       int
	Email    string
	Address  string
	Password string
}

type UserCreateReq struct {
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type UserLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//request payload validation
func (r *UserCreateReq) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Address, validation.Required),
		validation.Field(&r.Password, validation.Required))
}

func (r *UserLoginReq) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.Password, validation.Required))
}
