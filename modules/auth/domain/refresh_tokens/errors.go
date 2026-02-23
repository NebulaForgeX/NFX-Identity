package refresh_tokens

import "nfxid/pkgs/errx"

var (
	ErrRefreshTokenNotFound  = errx.NotFound("REFRESH_TOKEN_NOT_FOUND", "refresh token not found")
	ErrTokenIDRequired       = errx.InvalidArg("TOKEN_ID_REQUIRED", "token id is required")
	ErrUserIDRequired        = errx.InvalidArg("USER_ID_REQUIRED", "user id is required")
	ErrExpiresAtRequired     = errx.InvalidArg("EXPIRES_AT_REQUIRED", "expires at is required")
	ErrTokenIDAlreadyExists  = errx.Conflict("TOKEN_ID_ALREADY_EXISTS", "token id already exists")
	ErrTokenAlreadyRevoked   = errx.Conflict("TOKEN_ALREADY_REVOKED", "token already revoked")
	ErrTokenExpired          = errx.Expired("TOKEN_EXPIRED", "token expired")
	ErrInvalidRevokeReason   = errx.InvalidArg("INVALID_REVOKE_REASON", "invalid revoke reason")
)
