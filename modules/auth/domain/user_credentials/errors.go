package user_credentials

import "nfxid/pkgs/errx"

var (
	ErrUserCredentialNotFound     = errx.NotFound("USER_CREDENTIAL_NOT_FOUND", "user credential not found")
	ErrUserIDRequired             = errx.InvalidArg("USER_ID_REQUIRED", "user id is required")
	ErrCredentialTypeRequired     = errx.InvalidArg("CREDENTIAL_TYPE_REQUIRED", "credential type is required")
	ErrInvalidCredentialType      = errx.InvalidArg("INVALID_CREDENTIAL_TYPE", "invalid credential type")
	ErrInvalidCredentialStatus    = errx.InvalidArg("INVALID_CREDENTIAL_STATUS", "invalid credential status")
	ErrPasswordHashRequired       = errx.InvalidArg("PASSWORD_HASH_REQUIRED", "password hash is required")
	ErrUserCredentialAlreadyExists = errx.Conflict("USER_CREDENTIAL_ALREADY_EXISTS", "user credential already exists")
)
