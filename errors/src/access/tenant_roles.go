package access

import "nfxidentity/pkgs/errx"

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

/*
!TENANT_ROLE_NOT_FOUND
*en<tenant role not found>
*zh<租户角色不存在>
*fr<rôle tenant introuvable>

!TENANT_ROLE_KEY_REQUIRED
*en<tenant role key required>
*zh<租户角色键为必填>
*fr<clé de rôle tenant requise>

!TENANT_ID_REQUIRED
*en<tenant id required>
*zh<租户 ID 为必填>
*fr<id tenant requis>

!TENANT_ROLE_KEY_EXISTS_IN_TENANT
*en<tenant role key exists in tenant>
*zh<租户中已存在该角色键>
*fr<clé de rôle tenant existe déjà dans le tenant>

*/
