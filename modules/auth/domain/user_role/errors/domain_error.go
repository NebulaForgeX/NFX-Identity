package errors

import "errors"

var (
	ErrUserRoleNotFound     = errors.New("user role not found")
	ErrUserRoleAlreadyExists = errors.New("user role already exists")
	ErrUserIDRequired       = errors.New("user ID is required for user role")
	ErrRoleIDRequired       = errors.New("role ID is required for user role")
)

