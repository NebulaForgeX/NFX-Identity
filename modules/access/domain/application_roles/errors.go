package application_roles

import "errors"

var (
	ErrApplicationRoleNotFound            = errors.New("application role not found")
	ErrApplicationRoleApplicationIDRequired = errors.New("application_id is required")
	ErrApplicationRoleRoleKeyRequired     = errors.New("role_key is required")
)
