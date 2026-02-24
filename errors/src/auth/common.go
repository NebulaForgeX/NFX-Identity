package auth

import "nfxid/pkgs/errx"

// Shared codes and errors used by multiple auth domains.
const (
	CodeUserIDRequired      = "USER_ID_REQUIRED"
	CodeTenantIDRequired    = "TENANT_ID_REQUIRED"
	CodeExpiresAtRequired   = "EXPIRES_AT_REQUIRED"
	CodeInvalidRevokeReason = "INVALID_REVOKE_REASON"
)

var (
	ErrUserIDRequired      = errx.InvalidArg(CodeUserIDRequired, "user id is required")
	ErrTenantIDRequired    = errx.InvalidArg(CodeTenantIDRequired, "tenant id is required")
	ErrExpiresAtRequired   = errx.InvalidArg(CodeExpiresAtRequired, "expires at is required")
	ErrInvalidRevokeReason = errx.InvalidArg(CodeInvalidRevokeReason, "invalid revoke reason")
)
