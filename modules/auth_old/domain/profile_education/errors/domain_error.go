package errors

import "errors"

var (
	ErrEducationNotFound        = errors.New("education not found")
	ErrEducationSchoolRequired  = errors.New("education school is required")
	ErrEducationProfileIDRequired = errors.New("education profile_id is required")
)

