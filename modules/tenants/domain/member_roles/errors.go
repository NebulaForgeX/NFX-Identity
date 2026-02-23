package member_roles

import "nfxid/pkgs/errx"

var (
	ErrMemberRoleNotFound     = errx.NotFound("MEMBER_ROLE_NOT_FOUND", "member role not found")
	ErrTenantIDRequired       = errx.InvalidArg("TENANT_ID_REQUIRED", "tenant id is required")
	ErrMemberIDRequired       = errx.InvalidArg("MEMBER_ID_REQUIRED", "member id is required")
	ErrRoleIDRequired         = errx.InvalidArg("ROLE_ID_REQUIRED", "role id is required")
	ErrMemberRoleAlreadyExists = errx.Conflict("MEMBER_ROLE_ALREADY_EXISTS", "member role already exists")
	ErrMemberRoleAlreadyRevoked = errx.Conflict("MEMBER_ROLE_ALREADY_REVOKED", "member role already revoked")
	ErrMemberRoleExpired      = errx.Expired("MEMBER_ROLE_EXPIRED", "member role expired")
)
