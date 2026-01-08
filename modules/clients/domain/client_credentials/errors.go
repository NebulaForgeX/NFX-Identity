package client_credentials

import "errors"

var (
	ErrClientCredentialNotFound = errors.New("client credential not found")
	ErrClientIDRequired         = errors.New("client id is required")
	ErrAppIDRequired            = errors.New("app id is required")
	ErrSecretHashRequired       = errors.New("secret hash is required")
	ErrHashAlgRequired          = errors.New("hash alg is required")
	ErrClientIDAlreadyExists    = errors.New("client id already exists")
	ErrInvalidCredentialStatus  = errors.New("invalid credential status")
	ErrCredentialAlreadyRevoked = errors.New("credential already revoked")
	ErrCredentialExpired        = errors.New("credential expired")
)
