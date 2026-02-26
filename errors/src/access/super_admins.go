package access

import "nfxidentity/pkgs/errx"

const (
	CodeSuperAdminNotFound = "SUPER_ADMIN_NOT_FOUND"
)

var (
	ErrSuperAdminNotFound = errx.NotFound(CodeSuperAdminNotFound, "super admin not found")
)

/*
!SUPER_ADMIN_NOT_FOUND
*en<super admin not found>
*zh<超级管理员不存在>
*fr<super admin introuvable>

*/
