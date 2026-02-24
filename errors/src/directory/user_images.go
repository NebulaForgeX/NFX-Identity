package directory

import "nfxid/pkgs/errx"

const (
	CodeUserImageNotFound = "USER_IMAGE_NOT_FOUND"
)

var (
	ErrUserImageNotFound = errx.NotFound(CodeUserImageNotFound, "user image not found")
)

/*
!USER_IMAGE_NOT_FOUND
*en<user image not found>
*zh<用户图片不存在>
*fr<image utilisateur introuvable>

*/
