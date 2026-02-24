package tenants

import "nfxid/pkgs/errx"

const (
	CodeGroupNotFound        = "GROUP_NOT_FOUND"
	CodeGroupIDAlreadyExists = "GROUP_ID_ALREADY_EXISTS"
	CodeInvalidGroupType     = "INVALID_GROUP_TYPE"
)

var (
	ErrGroupNotFound        = errx.NotFound(CodeGroupNotFound, "group not found")
	ErrGroupIDAlreadyExists = errx.Conflict(CodeGroupIDAlreadyExists, "group id already exists")
	ErrInvalidGroupType     = errx.InvalidArg(CodeInvalidGroupType, "invalid group type")
)
