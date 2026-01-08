package scopes

import "errors"

var (
	ErrScopeNotFound      = errors.New("scope not found")
	ErrScopeRequired      = errors.New("scope is required")
	ErrScopeAlreadyExists = errors.New("scope already exists")
	ErrSystemScopeDelete  = errors.New("cannot delete system scope")
)
