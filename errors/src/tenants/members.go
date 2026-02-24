package tenants

import "nfxid/pkgs/errx"

const (
	CodeMemberNotFound      = "MEMBER_NOT_FOUND"
	CodeUserIDRequired      = "USER_ID_REQUIRED"
	CodeMemberAlreadyExists = "MEMBER_ALREADY_EXISTS"
	CodeInvalidMemberStatus = "INVALID_MEMBER_STATUS"
	CodeInvalidMemberSource = "INVALID_MEMBER_SOURCE"
)

var (
	ErrMemberNotFound      = errx.NotFound(CodeMemberNotFound, "member not found")
	ErrMemberAlreadyExists = errx.Conflict(CodeMemberAlreadyExists, "member already exists")
	ErrUserIDRequired      = errx.InvalidArg(CodeUserIDRequired, "user id is required")
	ErrInvalidMemberStatus = errx.InvalidArg(CodeInvalidMemberStatus, "invalid member status")
	ErrInvalidMemberSource = errx.InvalidArg(CodeInvalidMemberSource, "invalid member source")
)
