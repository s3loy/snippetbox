package models

import (
	"errors"
)

var (
	ErrNoRecord           = errors.New("models: no matching recort found")
	ErrInvalidCredentials = errors.New("models: indalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)
