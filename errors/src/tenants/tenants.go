package tenants

import "nfxid/pkgs/errx"

const (
	CodeTenantNotFound        = "TENANT_NOT_FOUND"
	CodeTenantIDAlreadyExists = "TENANT_ID_ALREADY_EXISTS"
	CodeInvalidTenantStatus   = "INVALID_TENANT_STATUS"
)

var (
	ErrTenantNotFound        = errx.NotFound(CodeTenantNotFound, "tenant not found")
	ErrTenantIDAlreadyExists = errx.Conflict(CodeTenantIDAlreadyExists, "tenant id already exists")
	ErrInvalidTenantStatus   = errx.InvalidArg(CodeInvalidTenantStatus, "invalid tenant status")
)
