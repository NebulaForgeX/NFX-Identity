package directory

import "nfxid/pkgs/errx"

const (
	CodeBadgeNotFound     = "BADGE_NOT_FOUND"
	CodeNameRequired      = "NAME_REQUIRED"
	CodeNameAlreadyExists = "NAME_ALREADY_EXISTS"
)

var (
	ErrBadgeNotFound     = errx.NotFound(CodeBadgeNotFound, "badge not found")
	ErrNameRequired      = errx.InvalidArg(CodeNameRequired, "name is required")
	ErrNameAlreadyExists = errx.Conflict(CodeNameAlreadyExists, "name already exists")
)
