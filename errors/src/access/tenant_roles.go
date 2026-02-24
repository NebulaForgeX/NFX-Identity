package access

import "nfxid/pkgs/errx"

const (
	CodeTenantRoleNotFound          = "TENANT_ROLE_NOT_FOUND"
	CodeTenantRoleKeyRequired       = "TENANT_ROLE_KEY_REQUIRED"
	CodeTenantIDRequired            = "TENANT_ID_REQUIRED"
	CodeTenantRoleKeyExistsInTenant = "TENANT_ROLE_KEY_EXISTS_IN_TENANT"
)

var (
	ErrTenantRoleNotFound          = errx.NotFound(CodeTenantRoleNotFound, "tenant role not found")
	ErrTenantRoleRoleKeyRequired   = errx.InvalidArg(CodeTenantRoleKeyRequired, "tenant role key is required")
	ErrTenantRoleTenantIDRequired  = errx.InvalidArg(CodeTenantIDRequired, "tenant id is required")
	ErrTenantRoleKeyExistsInTenant = errx.Conflict(CodeTenantRoleKeyExistsInTenant, "tenant role key already exists in tenant")
)
