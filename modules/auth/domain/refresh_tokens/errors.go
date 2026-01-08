package refresh_tokens

import "errors"

var (
	ErrRefreshTokenNotFound     = errors.New("refresh token not found")
	ErrTokenIDRequired          = errors.New("token id is required")
	ErrUserIDRequired           = errors.New("user id is required")
	ErrTenantIDRequired         = errors.New("tenant id is required")
	ErrExpiresAtRequired        = errors.New("expires at is required")
	ErrTokenIDAlreadyExists     = errors.New("token id already exists")
	ErrTokenAlreadyRevoked      = errors.New("token already revoked")
	ErrTokenExpired             = errors.New("token expired")
	ErrInvalidRevokeReason      = errors.New("invalid revoke reason")
)
