package user_educations

import "errors"

var (
	ErrUserEducationNotFound = errors.New("user education not found")
	ErrUserIDRequired        = errors.New("user id is required")
	ErrSchoolRequired        = errors.New("school is required")
)
