package mfa_factors

import "errors"

var (
	ErrMFAFactorNotFound      = errors.New("mfa factor not found")
	ErrFactorIDRequired       = errors.New("factor id is required")
	ErrUserIDRequired         = errors.New("user id is required")
	ErrTenantIDRequired       = errors.New("tenant id is required")
	ErrTypeRequired           = errors.New("type is required")
	ErrFactorIDAlreadyExists  = errors.New("factor id already exists")
	ErrInvalidMFAType         = errors.New("invalid mfa type")
)
