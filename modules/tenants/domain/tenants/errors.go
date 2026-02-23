package tenants

import "nfxid/pkgs/errx"

var (
	ErrTenantNotFound       = errx.NotFound("TENANT_NOT_FOUND", "tenant not found")
	ErrTenantIDRequired     = errx.InvalidArg("TENANT_ID_REQUIRED", "tenant id is required")
	ErrNameRequired         = errx.InvalidArg("NAME_REQUIRED", "name is required")
	ErrTenantIDAlreadyExists = errx.Conflict("TENANT_ID_ALREADY_EXISTS", "tenant id already exists")
	ErrInvalidTenantStatus  = errx.InvalidArg("INVALID_TENANT_STATUS", "invalid tenant status")
)
