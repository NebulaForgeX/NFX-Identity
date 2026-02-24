package auth

import "nfxid/pkgs/errx"

const (
	CodeMFAFactorNotFound     = "MFA_FACTOR_NOT_FOUND"
	CodeFactorIDRequired      = "FACTOR_ID_REQUIRED"
	CodeTypeRequired          = "TYPE_REQUIRED"
	CodeFactorIDAlreadyExists = "FACTOR_ID_ALREADY_EXISTS"
	CodeInvalidMFAType        = "INVALID_MFA_TYPE"
)

var (
	ErrMFAFactorNotFound     = errx.NotFound(CodeMFAFactorNotFound, "mfa factor not found")
	ErrFactorIDRequired      = errx.InvalidArg(CodeFactorIDRequired, "factor id is required")
	ErrTypeRequired          = errx.InvalidArg(CodeTypeRequired, "type is required")
	ErrFactorIDAlreadyExists = errx.Conflict(CodeFactorIDAlreadyExists, "factor id already exists")
	ErrInvalidMFAType        = errx.InvalidArg(CodeInvalidMFAType, "invalid mfa type")
)
