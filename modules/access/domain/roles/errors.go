package roles

import "errors"

var (
	ErrRoleNotFound      = errors.New("role not found")
	ErrRoleKeyRequired   = errors.New("role key is required")
	ErrRoleNameRequired  = errors.New("role name is required")
	ErrRoleKeyExists     = errors.New("role key already exists")
	ErrSystemRoleDelete  = errors.New("cannot delete system role")
	ErrInvalidScopeType  = errors.New("invalid scope type")
)
