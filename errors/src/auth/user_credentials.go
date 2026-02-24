package auth

import "nfxid/pkgs/errx"

const (
	CodeUserCredentialNotFound      = "USER_CREDENTIAL_NOT_FOUND"
	CodeCredentialTypeRequired      = "CREDENTIAL_TYPE_REQUIRED"
	CodeInvalidCredentialType       = "INVALID_CREDENTIAL_TYPE"
	CodeInvalidCredentialStatus     = "INVALID_CREDENTIAL_STATUS"
	CodePasswordHashRequired        = "PASSWORD_HASH_REQUIRED"
	CodeUserCredentialAlreadyExists = "USER_CREDENTIAL_ALREADY_EXISTS"
)

var (
	ErrUserCredentialNotFound      = errx.NotFound(CodeUserCredentialNotFound, "user credential not found")
	ErrCredentialTypeRequired      = errx.InvalidArg(CodeCredentialTypeRequired, "credential type is required")
	ErrInvalidCredentialType       = errx.InvalidArg(CodeInvalidCredentialType, "invalid credential type")
	ErrInvalidCredentialStatus     = errx.InvalidArg(CodeInvalidCredentialStatus, "invalid credential status")
	ErrPasswordHashRequired        = errx.InvalidArg(CodePasswordHashRequired, "password hash is required")
	ErrUserCredentialAlreadyExists = errx.Conflict(CodeUserCredentialAlreadyExists, "user credential already exists")
)
