package errors

import "errors"

var (
	ErrAuthorizationCodeNotFound       = errors.New("authorization code not found")
	ErrAuthorizationCodeCodeRequired   = errors.New("authorization code is required")
	ErrAuthorizationCodeCodeInvalid    = errors.New("authorization code is invalid")
	ErrAuthorizationCodeMaxUsesInvalid = errors.New("authorization code max uses must be greater than 0")
	ErrAuthorizationCodeAlreadyUsed    = errors.New("authorization code has reached maximum uses")
	ErrAuthorizationCodeExpired        = errors.New("authorization code has expired")
	ErrAuthorizationCodeInactive       = errors.New("authorization code is inactive")
)
