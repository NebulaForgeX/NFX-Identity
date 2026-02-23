package api_keys

import "nfxid/pkgs/errx"

var (
	ErrAPIKeyNotFound       = errx.NotFound("API_KEY_NOT_FOUND", "api key not found")
	ErrKeyIDRequired        = errx.InvalidArg("KEY_ID_REQUIRED", "key id is required")
	ErrAppIDRequired        = errx.InvalidArg("APP_ID_REQUIRED", "app id is required")
	ErrKeyHashRequired      = errx.InvalidArg("KEY_HASH_REQUIRED", "key hash is required")
	ErrHashAlgRequired      = errx.InvalidArg("HASH_ALG_REQUIRED", "hash alg is required")
	ErrNameRequired         = errx.InvalidArg("NAME_REQUIRED", "name is required")
	ErrKeyIDAlreadyExists   = errx.Conflict("KEY_ID_ALREADY_EXISTS", "key id already exists")
	ErrInvalidAPIKeyStatus  = errx.InvalidArg("INVALID_API_KEY_STATUS", "invalid api key status")
	ErrAPIKeyAlreadyRevoked = errx.Conflict("API_KEY_ALREADY_REVOKED", "api key already revoked")
	ErrAPIKeyExpired        = errx.Expired("API_KEY_EXPIRED", "api key expired")
)
