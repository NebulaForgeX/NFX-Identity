package super_admins

import "nfxid/pkgs/errx"

var (
	ErrSuperAdminNotFound = errx.NotFound("SUPER_ADMIN_NOT_FOUND", "super admin not found")
)
