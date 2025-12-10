package errors

import "errors"

var (
	ErrRoleNotFound      = errors.New("role not found")
	ErrRoleNameRequired  = errors.New("role name is required")
	ErrRoleAlreadyExists = errors.New("role already exists")
	ErrRoleNameExists    = errors.New("role name already exists")
	ErrSystemRoleDelete  = errors.New("cannot delete system role")
)

