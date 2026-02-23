package sessions

import "nfxid/pkgs/errx"

var (
	ErrSessionNotFound        = errx.NotFound("SESSION_NOT_FOUND", "session not found")
	ErrSessionIDRequired      = errx.InvalidArg("SESSION_ID_REQUIRED", "session id is required")
	ErrUserIDRequired         = errx.InvalidArg("USER_ID_REQUIRED", "user id is required")
	ErrTenantIDRequired       = errx.InvalidArg("TENANT_ID_REQUIRED", "tenant id is required")
	ErrExpiresAtRequired      = errx.InvalidArg("EXPIRES_AT_REQUIRED", "expires at is required")
	ErrSessionIDAlreadyExists = errx.Conflict("SESSION_ID_ALREADY_EXISTS", "session id already exists")
	ErrSessionAlreadyRevoked  = errx.Conflict("SESSION_ALREADY_REVOKED", "session already revoked")
	ErrSessionExpired         = errx.Expired("SESSION_EXPIRED", "session expired")
	ErrInvalidRevokeReason    = errx.InvalidArg("INVALID_REVOKE_REASON", "invalid revoke reason")
)
