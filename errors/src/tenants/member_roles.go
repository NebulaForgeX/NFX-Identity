package tenants

import "nfxid/pkgs/errx"

const (
	CodeMemberRoleNotFound       = "MEMBER_ROLE_NOT_FOUND"
	CodeMemberRoleAlreadyExists  = "MEMBER_ROLE_ALREADY_EXISTS"
	CodeMemberRoleAlreadyRevoked = "MEMBER_ROLE_ALREADY_REVOKED"
	CodeMemberRoleExpired        = "MEMBER_ROLE_EXPIRED"
)

var (
	ErrMemberRoleNotFound       = errx.NotFound(CodeMemberRoleNotFound, "member role not found")
	ErrMemberRoleAlreadyExists  = errx.Conflict(CodeMemberRoleAlreadyExists, "member role already exists")
	ErrMemberRoleAlreadyRevoked = errx.Conflict(CodeMemberRoleAlreadyRevoked, "member role already revoked")
	ErrMemberRoleExpired        = errx.Expired(CodeMemberRoleExpired, "member role expired")
)
