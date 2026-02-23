package tenant_apps

import "nfxid/pkgs/errx"

var (
	ErrTenantAppNotFound      = errx.NotFound("TENANT_APP_NOT_FOUND", "tenant app not found")
	ErrTenantIDRequired       = errx.InvalidArg("TENANT_ID_REQUIRED", "tenant id is required")
	ErrAppIDRequired          = errx.InvalidArg("APP_ID_REQUIRED", "app id is required")
	ErrTenantAppAlreadyExists = errx.Conflict("TENANT_APP_ALREADY_EXISTS", "tenant app already exists")
	ErrInvalidTenantAppStatus = errx.InvalidArg("INVALID_TENANT_APP_STATUS", "invalid tenant app status")
)
