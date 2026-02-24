package tenants

import "nfxid/pkgs/errx"

const (
	CodeMemberGroupNotFound       = "MEMBER_GROUP_NOT_FOUND"
	CodeMemberGroupAlreadyExists  = "MEMBER_GROUP_ALREADY_EXISTS"
	CodeMemberGroupAlreadyRevoked = "MEMBER_GROUP_ALREADY_REVOKED"
)

var (
	ErrMemberGroupNotFound       = errx.NotFound(CodeMemberGroupNotFound, "member group not found")
	ErrMemberGroupAlreadyExists  = errx.Conflict(CodeMemberGroupAlreadyExists, "member group already exists")
	ErrMemberGroupAlreadyRevoked = errx.Conflict(CodeMemberGroupAlreadyRevoked, "member group already revoked")
)
