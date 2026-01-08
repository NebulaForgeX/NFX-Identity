package sessions

import "errors"

var (
	ErrSessionNotFound      = errors.New("session not found")
	ErrSessionIDRequired     = errors.New("session id is required")
	ErrUserIDRequired        = errors.New("user id is required")
	ErrTenantIDRequired      = errors.New("tenant id is required")
	ErrExpiresAtRequired     = errors.New("expires at is required")
	ErrSessionIDAlreadyExists = errors.New("session id already exists")
	ErrSessionAlreadyRevoked = errors.New("session already revoked")
	ErrSessionExpired        = errors.New("session expired")
	ErrInvalidRevokeReason   = errors.New("invalid revoke reason")
)
