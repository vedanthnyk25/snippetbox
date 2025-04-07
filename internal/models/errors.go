package models

import(
	"errors"
)

var(
	ErrNoRecord= errors.New("models: No matching record found!")

	ErrInvalidCredentials= errors.New("models: Invalid credentials")
	ErrDuplicateEmail= errors.New("models: Duplicate email")
)