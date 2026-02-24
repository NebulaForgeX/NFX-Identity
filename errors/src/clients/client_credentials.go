package clients

import "nfxid/pkgs/errx"

const (
	CodeClientCredentialNotFound = "CLIENT_CREDENTIAL_NOT_FOUND"
	CodeClientIDRequired         = "CLIENT_ID_REQUIRED"
	CodeSecretHashRequired       = "SECRET_HASH_REQUIRED"
	CodeClientIDAlreadyExists    = "CLIENT_ID_ALREADY_EXISTS"
	CodeInvalidCredentialStatus  = "INVALID_CREDENTIAL_STATUS"
	CodeCredentialAlreadyRevoked = "CREDENTIAL_ALREADY_REVOKED"
	CodeCredentialExpired        = "CREDENTIAL_EXPIRED"
)

var (
	ErrClientCredentialNotFound = errx.NotFound(CodeClientCredentialNotFound, "client credential not found")
	ErrClientIDRequired         = errx.InvalidArg(CodeClientIDRequired, "client id is required")
	ErrSecretHashRequired       = errx.InvalidArg(CodeSecretHashRequired, "secret hash is required")
	ErrClientIDAlreadyExists    = errx.Conflict(CodeClientIDAlreadyExists, "client id already exists")
	ErrInvalidCredentialStatus  = errx.InvalidArg(CodeInvalidCredentialStatus, "invalid credential status")
	ErrCredentialAlreadyRevoked = errx.Conflict(CodeCredentialAlreadyRevoked, "credential already revoked")
	ErrCredentialExpired        = errx.Expired(CodeCredentialExpired, "credential expired")
)
