package tenants

import "nfxidentity/pkgs/errx"

// Shared codes and errors used by multiple tenant domains.
const (
	CodeTenantIDRequired  = "TENANT_ID_REQUIRED"
	CodeAppIDRequired     = "APP_ID_REQUIRED"
	CodeNameRequired      = "NAME_REQUIRED"
	CodeMemberIDRequired  = "MEMBER_ID_REQUIRED"
	CodeRoleIDRequired    = "ROLE_ID_REQUIRED"
	CodeGroupIDRequired   = "GROUP_ID_REQUIRED"
	CodeExpiresAtRequired = "EXPIRES_AT_REQUIRED"
)

var (
	ErrTenantIDRequired  = errx.InvalidArg(CodeTenantIDRequired, "tenant id is required")
	ErrAppIDRequired     = errx.InvalidArg(CodeAppIDRequired, "app id is required")
	ErrNameRequired      = errx.InvalidArg(CodeNameRequired, "name is required")
	ErrMemberIDRequired  = errx.InvalidArg(CodeMemberIDRequired, "member id is required")
	ErrRoleIDRequired    = errx.InvalidArg(CodeRoleIDRequired, "role id is required")
	ErrGroupIDRequired   = errx.InvalidArg(CodeGroupIDRequired, "group id is required")
	ErrExpiresAtRequired = errx.InvalidArg(CodeExpiresAtRequired, "expires at is required")
)

/*
!TENANT_ID_REQUIRED
*en<tenant id required>
*zh<租户 ID 为必填>
*fr<id tenant requis>

!APP_ID_REQUIRED
*en<app id required>
*zh<应用 ID 为必填>
*fr<id d'application requis>

!NAME_REQUIRED
*en<name required>
*zh<名称为必填>
*fr<nom requis>

!MEMBER_ID_REQUIRED
*en<member id required>
*zh<成员 ID 为必填>
*fr<id de membre requis>

!ROLE_ID_REQUIRED
*en<role id required>
*zh<角色 ID 为必填>
*fr<id de rôle requis>

!GROUP_ID_REQUIRED
*en<group id required>
*zh<组 ID 为必填>
*fr<id de groupe requis>

!EXPIRES_AT_REQUIRED
*en<expires at required>
*zh<过期时间为必填>
*fr<date d'expiration requise>

*/
