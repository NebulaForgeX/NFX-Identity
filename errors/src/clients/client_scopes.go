package clients

import "nfxid/pkgs/errx"

const (
	CodeClientScopeNotFound       = "CLIENT_SCOPE_NOT_FOUND"
	CodeScopeRequired             = "SCOPE_REQUIRED"
	CodeClientScopeAlreadyExists  = "CLIENT_SCOPE_ALREADY_EXISTS"
	CodeClientScopeAlreadyRevoked = "CLIENT_SCOPE_ALREADY_REVOKED"
	CodeClientScopeExpired        = "CLIENT_SCOPE_EXPIRED"
)

var (
	ErrClientScopeNotFound       = errx.NotFound(CodeClientScopeNotFound, "client scope not found")
	ErrScopeRequired             = errx.InvalidArg(CodeScopeRequired, "scope is required")
	ErrClientScopeAlreadyExists  = errx.Conflict(CodeClientScopeAlreadyExists, "client scope already exists")
	ErrClientScopeAlreadyRevoked = errx.Conflict(CodeClientScopeAlreadyRevoked, "client scope already revoked")
	ErrClientScopeExpired        = errx.Expired(CodeClientScopeExpired, "client scope expired")
)
