package clients

import "nfxid/pkgs/errx"

const (
	CodeAPIKeyNotFound       = "API_KEY_NOT_FOUND"
	CodeKeyIDRequired        = "KEY_ID_REQUIRED"
	CodeKeyHashRequired      = "KEY_HASH_REQUIRED"
	CodeKeyIDAlreadyExists   = "KEY_ID_ALREADY_EXISTS"
	CodeInvalidAPIKeyStatus  = "INVALID_API_KEY_STATUS"
	CodeAPIKeyAlreadyRevoked = "API_KEY_ALREADY_REVOKED"
	CodeAPIKeyExpired        = "API_KEY_EXPIRED"
)

var (
	ErrAPIKeyNotFound       = errx.NotFound(CodeAPIKeyNotFound, "api key not found")
	ErrKeyIDRequired        = errx.InvalidArg(CodeKeyIDRequired, "key id is required")
	ErrKeyHashRequired      = errx.InvalidArg(CodeKeyHashRequired, "key hash is required")
	ErrKeyIDAlreadyExists   = errx.Conflict(CodeKeyIDAlreadyExists, "key id already exists")
	ErrInvalidAPIKeyStatus  = errx.InvalidArg(CodeInvalidAPIKeyStatus, "invalid api key status")
	ErrAPIKeyAlreadyRevoked = errx.Conflict(CodeAPIKeyAlreadyRevoked, "api key already revoked")
	ErrAPIKeyExpired        = errx.Expired(CodeAPIKeyExpired, "api key expired")
)
