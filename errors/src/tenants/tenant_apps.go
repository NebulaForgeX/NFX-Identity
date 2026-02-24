package tenants

import "nfxid/pkgs/errx"

const (
	CodeTenantAppNotFound      = "TENANT_APP_NOT_FOUND"
	CodeTenantAppAlreadyExists = "TENANT_APP_ALREADY_EXISTS"
	CodeInvalidTenantAppStatus = "INVALID_TENANT_APP_STATUS"
)

var (
	ErrTenantAppNotFound      = errx.NotFound(CodeTenantAppNotFound, "tenant app not found")
	ErrTenantAppAlreadyExists = errx.Conflict(CodeTenantAppAlreadyExists, "tenant app already exists")
	ErrInvalidTenantAppStatus = errx.InvalidArg(CodeInvalidTenantAppStatus, "invalid tenant app status")
)
