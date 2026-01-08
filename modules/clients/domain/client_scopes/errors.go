package client_scopes

import "errors"

var (
	ErrClientScopeNotFound     = errors.New("client scope not found")
	ErrAppIDRequired           = errors.New("app id is required")
	ErrScopeRequired           = errors.New("scope is required")
	ErrClientScopeAlreadyExists = errors.New("client scope already exists")
	ErrClientScopeAlreadyRevoked = errors.New("client scope already revoked")
	ErrClientScopeExpired      = errors.New("client scope expired")
)
