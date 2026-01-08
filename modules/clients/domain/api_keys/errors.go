package api_keys

import "errors"

var (
	ErrAPIKeyNotFound      = errors.New("api key not found")
	ErrKeyIDRequired       = errors.New("key id is required")
	ErrAppIDRequired       = errors.New("app id is required")
	ErrKeyHashRequired     = errors.New("key hash is required")
	ErrHashAlgRequired     = errors.New("hash alg is required")
	ErrNameRequired        = errors.New("name is required")
	ErrKeyIDAlreadyExists  = errors.New("key id already exists")
	ErrInvalidAPIKeyStatus = errors.New("invalid api key status")
	ErrAPIKeyAlreadyRevoked = errors.New("api key already revoked")
	ErrAPIKeyExpired       = errors.New("api key expired")
)
