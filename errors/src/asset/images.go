package asset

import "nfxidentity/pkgs/errx"

// 栅格图（asset."Images"）相关
var (
	ErrImageNotFound           = errx.NotFound("IMAGE_NOT_FOUND", "image not found")
	ErrImageFilePathRequired   = errx.InvalidArg("IMAGE_FILE_PATH_REQUIRED", "image file path is required")
	ErrImageFileNameRequired   = errx.InvalidArg("IMAGE_FILE_NAME_REQUIRED", "image file name is required")
	ErrImageFileSizeInvalid    = errx.InvalidArg("IMAGE_FILE_SIZE_INVALID", "image file size must be greater than 0")
	ErrImageMimeTypeRequired   = errx.InvalidArg("IMAGE_MIME_TYPE_REQUIRED", "image mime type is required")
	ErrImageUploaderIDRequired = errx.InvalidArg("IMAGE_UPLOADER_ID_REQUIRED", "image uploader id is required")
	ErrImageAlreadyExists      = errx.Conflict("IMAGE_ALREADY_EXISTS", "image already exists")
	ErrImageMoveOnlyFromTmp    = errx.InvalidArg("IMAGE_MOVE_ONLY_FROM_TMP", "image can only be moved from tmp path")
	ErrImageDeleteNotTmp       = errx.InvalidArg("IMAGE_DELETE_NOT_TMP", "only tmp images can be deleted")
	ErrImageInvalidTargetType  = errx.InvalidArg("IMAGE_INVALID_TARGET_TYPE", "invalid image move target type")
	ErrInvalidImageID          = errx.InvalidArg("IMAGE_INVALID_ID", "invalid image id")
	ErrInvalidAccountID        = errx.InvalidArg("IMAGE_INVALID_ACCOUNT_ID", "invalid account id")
	ErrInvalidProfileID        = errx.InvalidArg("IMAGE_INVALID_PROFILE_ID", "invalid profile id")
	ErrInvalidAssetType        = errx.InvalidArg("IMAGE_INVALID_ASSET_TYPE", "invalid asset type")
	ErrInvalidProfileScope     = errx.InvalidArg("IMAGE_INVALID_PROFILE_SCOPE", "invalid profile scope")
	ErrInvalidAssetCategory    = errx.InvalidArg("IMAGE_INVALID_ASSET_CATEGORY", "invalid asset category")
	ErrImageRequestBodyInvalid = errx.InvalidArg("IMAGE_REQUEST_BODY_INVALID", "invalid image request body")
	ErrInvalidFile             = errx.InvalidArg("IMAGE_INVALID_FILE", "missing or invalid file")
	ErrMissingUploaderID       = errx.InvalidArg("IMAGE_MISSING_UPLOADER_ID", "missing uploader_id")
	ErrInvalidUploaderID       = errx.InvalidArg("IMAGE_INVALID_UPLOADER_ID", "invalid uploader_id")
	ErrInvalidTargetID         = errx.InvalidArg("IMAGE_INVALID_TARGET_ID", "invalid target_id")
	ErrImageUploadFailed       = errx.Internal("IMAGE_UPLOAD_FAILED", "image upload failed")
	ErrImageCreateFailed       = errx.Internal("IMAGE_CREATE_FAILED", "failed to create image")
	ErrImageUpdateFailed       = errx.Internal("IMAGE_UPDATE_FAILED", "failed to update image")
	ErrImageDeleteFailed       = errx.Internal("IMAGE_DELETE_FAILED", "failed to delete image")
	ErrImageQueryFailed        = errx.Internal("IMAGE_QUERY_FAILED", "image query failed")
)

/*
!IMAGE_NOT_FOUND
*en<Image not found>
*zh<图片不存在>
*fr<image introuvable>

!IMAGE_FILE_PATH_REQUIRED
*en<Image file path is required>
*zh<图片文件路径必填>
*fr<chemin du fichier image requis>

!IMAGE_FILE_NAME_REQUIRED
*en<Image file name is required>
*zh<图片文件名必填>
*fr<nom du fichier image requis>

!IMAGE_FILE_SIZE_INVALID
*en<Image file size must be greater than 0>
*zh<图片文件大小必须大于 0>
*fr<taille du fichier image doit être supérieure à 0>

!IMAGE_MIME_TYPE_REQUIRED
*en<Image mime type is required>
*zh<图片 MIME 类型必填>
*fr<type MIME de l'image requis>

!IMAGE_UPLOADER_ID_REQUIRED
*en<Image uploader id is required>
*zh<图片上传者 ID 必填>
*fr<ID de l'uploader requis>

!IMAGE_ALREADY_EXISTS
*en<Image already exists>
*zh<图片已存在>
*fr<image déjà existante>

!IMAGE_MOVE_ONLY_FROM_TMP
*en<Image can only be moved from tmp path>
*zh<仅支持从临时路径移动图片>
*fr<l'image ne peut être déplacée que depuis le chemin temporaire>

!IMAGE_DELETE_NOT_TMP
*en<Only tmp images can be deleted>
*zh<仅可删除临时图片>
*fr<seules les images temporaires peuvent être supprimées>

!IMAGE_INVALID_TARGET_TYPE
*en<Invalid target type: must be account, profile, or education>
*zh<无效目标类型：须为 account、profile 或 education>
*fr<type de cible invalide : doit être account, profile ou education>

!IMAGE_INVALID_ID
*en<Invalid image id>
*zh<无效的图片 ID>
*fr<ID d'image invalide>

!IMAGE_INVALID_ACCOUNT_ID
*en<Invalid account id>
*zh<无效的账号 ID>
*fr<ID de compte invalide>

!IMAGE_INVALID_PROFILE_ID
*en<Invalid profile id>
*zh<无效的 Profile ID>
*fr<ID de profil invalide>

!IMAGE_INVALID_ASSET_TYPE
*en<Invalid asset type>
*zh<无效的资源类型>
*fr<type de ressource invalide>

!IMAGE_INVALID_PROFILE_SCOPE
*en<Invalid profile scope>
*zh<无效的 profile scope>
*fr<profile scope invalide>

!IMAGE_INVALID_ASSET_CATEGORY
*en<Invalid asset category>
*zh<无效的资源 category>
*fr<catégorie de ressource invalide>

!IMAGE_REQUEST_BODY_INVALID
*en<Invalid image request body>
*zh<无效的图片请求体>
*fr<corps de requête image invalide>

!IMAGE_INVALID_FILE
*en<Missing or invalid file>
*zh<缺少或无效的文件>
*fr<fichier manquant ou invalide>

!IMAGE_MISSING_UPLOADER_ID
*en<Missing uploader_id>
*zh<缺少 uploader_id>
*fr<uploader_id manquant>

!IMAGE_INVALID_UPLOADER_ID
*en<Invalid uploader_id>
*zh<无效的 uploader_id>
*fr<uploader_id invalide>

!IMAGE_INVALID_TARGET_ID
*en<Invalid target_id>
*zh<无效的 target_id>
*fr<target_id invalide>

!IMAGE_UPLOAD_FAILED
*en<Image upload failed>
*zh<图片上传失败>
*fr<échec du téléchargement de l'image>
*p*en<Sorry, we couldn't upload your image right now. This is on us—please try again later.>
*p*zh<抱歉，图片暂时上传失败，这是我们的问题，请稍后再试。>
*p*fr<Désolé, l'image n'a pas pu être téléversée. Veuillez réessayer plus tard.>

!IMAGE_CREATE_FAILED
*en<Failed to create image>
*zh<创建图片记录失败>
*fr<échec de la création de l'image>
*p*en<Sorry, we couldn't save your image right now. Please try again later.>
*p*zh<抱歉，暂时无法保存图片，请稍后再试。>
*p*fr<Désolé, l'image n'a pas pu être enregistrée. Veuillez réessayer plus tard.>

!IMAGE_UPDATE_FAILED
*en<Failed to update image>
*zh<更新图片失败>
*fr<échec de la mise à jour de l'image>
*p*en<Sorry, we couldn't update your image right now. Please try again later.>
*p*zh<抱歉，暂时无法更新图片，请稍后再试。>
*p*fr<Désolé, l'image n'a pas pu être mise à jour. Veuillez réessayer plus tard.>

!IMAGE_DELETE_FAILED
*en<Failed to delete image>
*zh<删除图片失败>
*fr<échec de la suppression de l'image>
*p*en<Sorry, we couldn't delete this image right now. Please try again later.>
*p*zh<抱歉，暂时无法删除图片，请稍后再试。>
*p*fr<Désolé, l'image n'a pas pu être supprimée. Veuillez réessayer plus tard.>

!IMAGE_QUERY_FAILED
*en<Image query failed>
*zh<图片查询失败>
*fr<échec de la requête image>
*p*en<Sorry, we couldn't load this image right now. Please try again later.>
*p*zh<抱歉，暂时无法加载图片，请稍后再试。>
*p*fr<Désolé, l'image n'a pas pu être chargée. Veuillez réessayer plus tard.>
*/
