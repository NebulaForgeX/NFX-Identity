package directory

import "nfxidentity/pkgs/errx"

const (
	CodeUserAvatarNotFound = "USER_AVATAR_NOT_FOUND"
)

var (
	ErrUserAvatarNotFound = errx.NotFound(CodeUserAvatarNotFound, "user avatar not found")
)

/*
!USER_AVATAR_NOT_FOUND
*en<user avatar not found>
*zh<用户头像不存在>
*fr<avatar utilisateur introuvable>

*/
