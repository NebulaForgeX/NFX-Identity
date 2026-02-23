package tenant_roles

import "nfxid/pkgs/errx"

var (
	ErrTenantRoleNotFound          = errx.NotFound("TENANT_ROLE_NOT_FOUND", "tenant role not found")
	ErrTenantRoleRoleKeyRequired   = errx.InvalidArg("TENANT_ROLE_KEY_REQUIRED", "tenant role key is required")
	ErrTenantRoleTenantIDRequired  = errx.InvalidArg("TENANT_ID_REQUIRED", "tenant id is required")
	ErrTenantRoleKeyExistsInTenant = errx.Conflict("TENANT_ROLE_KEY_EXISTS_IN_TENANT", "tenant role key already exists in tenant")
)
