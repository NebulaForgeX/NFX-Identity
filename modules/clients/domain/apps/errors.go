package apps

import "nfxid/pkgs/errx"

var (
	ErrAppNotFound        = errx.NotFound("APP_NOT_FOUND", "app not found")
	ErrAppIDRequired      = errx.InvalidArg("APP_ID_REQUIRED", "app id is required")
	ErrNameRequired       = errx.InvalidArg("NAME_REQUIRED", "name is required")
	ErrTenantIDRequired   = errx.InvalidArg("TENANT_ID_REQUIRED", "tenant id is required")
	ErrAppIDAlreadyExists = errx.Conflict("APP_ID_ALREADY_EXISTS", "app id already exists")
	ErrInvalidAppType     = errx.InvalidArg("INVALID_APP_TYPE", "invalid app type")
	ErrInvalidAppStatus   = errx.InvalidArg("INVALID_APP_STATUS", "invalid app status")
	ErrInvalidEnvironment = errx.InvalidArg("INVALID_ENVIRONMENT", "invalid environment")
)
