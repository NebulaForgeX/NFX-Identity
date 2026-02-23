package client_scopes

import "nfxid/pkgs/errx"

var (
	ErrClientScopeNotFound      = errx.NotFound("CLIENT_SCOPE_NOT_FOUND", "client scope not found")
	ErrAppIDRequired            = errx.InvalidArg("APP_ID_REQUIRED", "app id is required")
	ErrScopeRequired            = errx.InvalidArg("SCOPE_REQUIRED", "scope is required")
	ErrClientScopeAlreadyExists  = errx.Conflict("CLIENT_SCOPE_ALREADY_EXISTS", "client scope already exists")
	ErrClientScopeAlreadyRevoked = errx.Conflict("CLIENT_SCOPE_ALREADY_REVOKED", "client scope already revoked")
	ErrClientScopeExpired       = errx.Expired("CLIENT_SCOPE_EXPIRED", "client scope expired")
)
