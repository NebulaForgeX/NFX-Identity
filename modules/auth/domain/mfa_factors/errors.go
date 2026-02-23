package mfa_factors

import "nfxid/pkgs/errx"

var (
	ErrMFAFactorNotFound     = errx.NotFound("MFA_FACTOR_NOT_FOUND", "mfa factor not found")
	ErrFactorIDRequired      = errx.InvalidArg("FACTOR_ID_REQUIRED", "factor id is required")
	ErrUserIDRequired        = errx.InvalidArg("USER_ID_REQUIRED", "user id is required")
	ErrTenantIDRequired      = errx.InvalidArg("TENANT_ID_REQUIRED", "tenant id is required")
	ErrTypeRequired          = errx.InvalidArg("TYPE_REQUIRED", "type is required")
	ErrFactorIDAlreadyExists = errx.Conflict("FACTOR_ID_ALREADY_EXISTS", "factor id already exists")
	ErrInvalidMFAType        = errx.InvalidArg("INVALID_MFA_TYPE", "invalid mfa type")
)
