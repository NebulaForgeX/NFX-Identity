package asset

import "nfxidentity/pkgs/errx"

var (
	ErrAssetNotFound    = errx.NotFound("ASSET_NOT_FOUND", "asset not found")
	ErrInvalidAssetKind = errx.InvalidArg("INVALID_ASSET_KIND", "kind must be images, files, videos, or audios")
	ErrPresignFailed    = errx.Internal("PRESIGN_FAILED", "failed to presign upload")
	ErrObjectMissing    = errx.FailedPrecond("OBJECT_MISSING", "upload not found in object storage")
)

/*
!ASSET_NOT_FOUND
*en<Asset not found>
*zh<资源不存在>
*fr<Ressource introuvable>

!INVALID_ASSET_KIND
*en<Kind must be images, files, videos, or audios>
*zh<类型必须是 images、files、videos 或 audios>
*fr<Le type doit être images, files, videos ou audios>

!PRESIGN_FAILED
*en<Failed to presign upload>
*zh<预签名上传失败>
*fr<Échec de la pré-signature d'upload>

!OBJECT_MISSING
*en<Upload not found in object storage>
*zh<对象存储中找不到上传文件>
*fr<Fichier introuvable dans le stockage objet>
*/
