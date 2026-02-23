package member_app_roles

import "nfxid/pkgs/errx"

var (
	ErrMemberAppRoleNotFound     = errx.NotFound("MEMBER_APP_ROLE_NOT_FOUND", "member app role not found")
	ErrMemberIDRequired          = errx.InvalidArg("MEMBER_ID_REQUIRED", "member id is required")
	ErrAppIDRequired             = errx.InvalidArg("APP_ID_REQUIRED", "app id is required")
	ErrRoleIDRequired            = errx.InvalidArg("ROLE_ID_REQUIRED", "role id is required")
	ErrMemberAppRoleAlreadyExists = errx.Conflict("MEMBER_APP_ROLE_ALREADY_EXISTS", "member app role already exists")
	ErrMemberAppRoleAlreadyRevoked = errx.Conflict("MEMBER_APP_ROLE_ALREADY_REVOKED", "member app role already revoked")
	ErrMemberAppRoleExpired      = errx.Expired("MEMBER_APP_ROLE_EXPIRED", "member app role expired")
)
