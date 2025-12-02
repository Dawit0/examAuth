package domain

import "errors"

var (
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password length")
	ErrInvalidUsername = errors.New("enter username")
	ErrInvalidPhone    = errors.New("invalid phone number")
)
