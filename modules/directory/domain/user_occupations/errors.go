package user_occupations

import "errors"

var (
	ErrUserOccupationNotFound = errors.New("user occupation not found")
	ErrUserIDRequired         = errors.New("user id is required")
	ErrCompanyRequired        = errors.New("company is required")
	ErrPositionRequired       = errors.New("position is required")
)
