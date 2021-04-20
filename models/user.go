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

//request model
type UserCreateReq struct {
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type UserLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateReq struct {
	ID       int    `json:"id"`
	Email    string `json:"email,omitempty"`
	Address  string `json:"address,omitempty"`
	Password string `json:"password,omitempty"`
}

//response model
type UserProfileResp struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

//request payload validation
func (r UserCreateReq) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Address, validation.Required),
		validation.Field(&r.Password, validation.Required))
}

func (r UserLoginReq) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.Password, validation.Required))
}

func (r UserUpdateReq) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.ID, validation.Required),
		validation.Field(&r.Email, is.Email))
}
