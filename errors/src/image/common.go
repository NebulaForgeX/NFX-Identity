package image

import "nfxidentity/pkgs/errx"

// Shared codes and errors used by multiple image domains.
const (
	CodeStoragePathRequired = "STORAGE_PATH_REQUIRED"
	CodeImageIDRequired     = "IMAGE_ID_REQUIRED"
)

var (
	ErrStoragePathRequired = errx.InvalidArg(CodeStoragePathRequired, "storage path is required")
	ErrImageIDRequired     = errx.InvalidArg(CodeImageIDRequired, "image id is required")
)

/*
!STORAGE_PATH_REQUIRED
*en<storage path required>
*zh<存储路径为必填>
*fr<chemin de stockage requis>

!IMAGE_ID_REQUIRED
*en<image id required>
*zh<图片 ID 为必填>
*fr<id d'image requis>

*/
