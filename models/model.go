package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type User struct {
	UUID    uuid.UUID `json:"uuid"`
	Name    string    `json:"name" validate:"required,alpha"`
	Surname string    `json:"surname" validate:"required,alpha"`
	Email   string    `json:"email" validate:"required,email"`
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
