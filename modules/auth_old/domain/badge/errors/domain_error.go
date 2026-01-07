package errors

import "errors"

var (
	ErrInvalidName        = errors.New("invalid name")
	ErrInvalidDescription = errors.New("invalid description")
	ErrInvalidIconURL     = errors.New("invalid icon url")
	ErrInvalidColor       = errors.New("invalid color")
	ErrInvalidCategory    = errors.New("invalid category")
	ErrNotFound           = errors.New("badge not found")
	ErrNameAlreadyExists  = errors.New("badge name already exists")
)

