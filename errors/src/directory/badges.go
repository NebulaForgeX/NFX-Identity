package directory

import "nfxidentity/pkgs/errx"

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

/*
!BADGE_NOT_FOUND
*en<badge not found>
*zh<徽章不存在>
*fr<badge introuvable>

!NAME_REQUIRED
*en<name required>
*zh<名称为必填>
*fr<nom requis>

!NAME_ALREADY_EXISTS
*en<name already exists>
*zh<名称已存在>
*fr<nom existe déjà>

*/
