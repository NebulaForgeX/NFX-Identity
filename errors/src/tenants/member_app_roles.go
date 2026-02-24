package tenants

import "nfxid/pkgs/errx"

const (
	CodeMemberAppRoleNotFound       = "MEMBER_APP_ROLE_NOT_FOUND"
	CodeMemberAppRoleAlreadyExists  = "MEMBER_APP_ROLE_ALREADY_EXISTS"
	CodeMemberAppRoleAlreadyRevoked = "MEMBER_APP_ROLE_ALREADY_REVOKED"
	CodeMemberAppRoleExpired        = "MEMBER_APP_ROLE_EXPIRED"
)

var (
	ErrMemberAppRoleNotFound       = errx.NotFound(CodeMemberAppRoleNotFound, "member app role not found")
	ErrMemberAppRoleAlreadyExists  = errx.Conflict(CodeMemberAppRoleAlreadyExists, "member app role already exists")
	ErrMemberAppRoleAlreadyRevoked = errx.Conflict(CodeMemberAppRoleAlreadyRevoked, "member app role already revoked")
	ErrMemberAppRoleExpired        = errx.Expired(CodeMemberAppRoleExpired, "member app role expired")
)
