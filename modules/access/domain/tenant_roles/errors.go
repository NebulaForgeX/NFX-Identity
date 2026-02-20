package tenant_roles

import "errors"

var (
	ErrTenantRoleNotFound           = errors.New("tenant role not found")
	ErrTenantRoleRoleKeyRequired    = errors.New("tenant role key is required")
	ErrTenantRoleTenantIDRequired   = errors.New("tenant id is required")
	ErrTenantRoleKeyExistsInTenant  = errors.New("tenant role key already exists in tenant")
)
