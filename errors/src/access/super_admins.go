package access

import "nfxid/pkgs/errx"

const (
	CodeSuperAdminNotFound = "SUPER_ADMIN_NOT_FOUND"
)

var (
	ErrSuperAdminNotFound = errx.NotFound(CodeSuperAdminNotFound, "super admin not found")
)
