package clients

import "nfxid/pkgs/errx"

const (
	CodeAppNotFound        = "APP_NOT_FOUND"
	CodeTenantIDRequired   = "TENANT_ID_REQUIRED"
	CodeAppIDAlreadyExists = "APP_ID_ALREADY_EXISTS"
	CodeInvalidAppType     = "INVALID_APP_TYPE"
	CodeInvalidAppStatus   = "INVALID_APP_STATUS"
	CodeInvalidEnvironment = "INVALID_ENVIRONMENT"
)

var (
	ErrAppNotFound        = errx.NotFound(CodeAppNotFound, "app not found")
	ErrTenantIDRequired   = errx.InvalidArg(CodeTenantIDRequired, "tenant id is required")
	ErrAppIDAlreadyExists = errx.Conflict(CodeAppIDAlreadyExists, "app id already exists")
	ErrInvalidAppType     = errx.InvalidArg(CodeInvalidAppType, "invalid app type")
	ErrInvalidAppStatus   = errx.InvalidArg(CodeInvalidAppStatus, "invalid app status")
	ErrInvalidEnvironment = errx.InvalidArg(CodeInvalidEnvironment, "invalid environment")
)
