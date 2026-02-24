package tenants

import "nfxid/pkgs/errx"

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
