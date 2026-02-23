package client_credentials

import "nfxid/pkgs/errx"

var (
	ErrClientCredentialNotFound = errx.NotFound("CLIENT_CREDENTIAL_NOT_FOUND", "client credential not found")
	ErrClientIDRequired         = errx.InvalidArg("CLIENT_ID_REQUIRED", "client id is required")
	ErrAppIDRequired            = errx.InvalidArg("APP_ID_REQUIRED", "app id is required")
	ErrSecretHashRequired       = errx.InvalidArg("SECRET_HASH_REQUIRED", "secret hash is required")
	ErrHashAlgRequired           = errx.InvalidArg("HASH_ALG_REQUIRED", "hash alg is required")
	ErrClientIDAlreadyExists    = errx.Conflict("CLIENT_ID_ALREADY_EXISTS", "client id already exists")
	ErrInvalidCredentialStatus  = errx.InvalidArg("INVALID_CREDENTIAL_STATUS", "invalid credential status")
	ErrCredentialAlreadyRevoked = errx.Conflict("CREDENTIAL_ALREADY_REVOKED", "credential already revoked")
	ErrCredentialExpired        = errx.Expired("CREDENTIAL_EXPIRED", "credential expired")
)
