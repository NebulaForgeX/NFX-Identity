package directory

import "nfxid/pkgs/errx"

// Shared codes and errors used by multiple directory domains.
const (
	CodeUserIDRequired  = "USER_ID_REQUIRED"
	CodeImageIDRequired = "IMAGE_ID_REQUIRED"
)

var (
	ErrUserIDRequired  = errx.InvalidArg(CodeUserIDRequired, "user id is required")
	ErrImageIDRequired = errx.InvalidArg(CodeImageIDRequired, "image id is required")
)

/*
!USER_ID_REQUIRED
*en<user id required>
*zh<用户 ID 为必填>
*fr<identifiant utilisateur requis>

!IMAGE_ID_REQUIRED
*en<image id required>
*zh<图片 ID 为必填>
*fr<id d'image requis>

*/
