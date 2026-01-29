package constants

import "errors"

// Common API-level sentinel errors. Use these at HTTP/gRPC boundaries for consistent English messages.
// Domain-specific errors remain in each module's domain/application errors.go.
var (
	// ErrBadRequest indicates invalid request parameters or body (HTTP 400).
	ErrBadRequest = errors.New("invalid request")

	// ErrUnauthorized indicates missing or invalid credentials (HTTP 401).
	ErrUnauthorized = errors.New("unauthorized")

	// ErrForbidden indicates the identity is valid but not allowed to perform the action (HTTP 403).
	ErrForbidden = errors.New("forbidden")

	// ErrNotFound indicates the requested resource does not exist (HTTP 404).
	ErrNotFound = errors.New("resource not found")

	// ErrConflict indicates a conflict with current state (e.g. duplicate key) (HTTP 409).
	ErrConflict = errors.New("conflict with current state")

	// ErrInternal indicates an unexpected server error (HTTP 500).
	ErrInternal = errors.New("internal server error")
)
