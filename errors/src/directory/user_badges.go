package directory

import "nfxid/pkgs/errx"

const (
	CodeUserBadgeNotFound      = "USER_BADGE_NOT_FOUND"
	CodeBadgeIDRequired        = "BADGE_ID_REQUIRED"
	CodeUserBadgeAlreadyExists = "USER_BADGE_ALREADY_EXISTS"
)

var (
	ErrUserBadgeNotFound      = errx.NotFound(CodeUserBadgeNotFound, "user badge not found")
	ErrBadgeIDRequired        = errx.InvalidArg(CodeBadgeIDRequired, "badge id is required")
	ErrUserBadgeAlreadyExists = errx.Conflict(CodeUserBadgeAlreadyExists, "user badge already exists")
)

/*
!USER_BADGE_NOT_FOUND
*en<user badge not found>
*zh<用户徽章不存在>
*fr<badge utilisateur introuvable>

!BADGE_ID_REQUIRED
*en<badge id required>
*zh<徽章 ID 为必填>
*fr<id de badge requis>

!USER_BADGE_ALREADY_EXISTS
*en<user badge already exists>
*zh<用户徽章已存在>
*fr<badge utilisateur existe déjà>

*/
