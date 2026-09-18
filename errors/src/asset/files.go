package asset

import "nfxidentity/pkgs/errx"

// 通用文件（asset."Files"）相关
var (
	ErrFileNotFound           = errx.NotFound("FILE_NOT_FOUND", "file not found")
	ErrFileFilePathRequired   = errx.InvalidArg("FILE_FILE_PATH_REQUIRED", "file path is required")
	ErrFileFileNameRequired   = errx.InvalidArg("FILE_FILE_NAME_REQUIRED", "file name is required")
	ErrFileFileSizeInvalid    = errx.InvalidArg("FILE_FILE_SIZE_INVALID", "file size must be greater than 0")
	ErrFileMimeTypeRequired   = errx.InvalidArg("FILE_MIME_TYPE_REQUIRED", "file mime type is required")
	ErrFileUploaderIDRequired = errx.InvalidArg("FILE_UPLOADER_ID_REQUIRED", "file uploader id is required")
	ErrFileAlreadyExists      = errx.Conflict("FILE_ALREADY_EXISTS", "file already exists")
	ErrFileMoveOnlyFromTmp    = errx.InvalidArg("FILE_MOVE_ONLY_FROM_TMP", "file can only be moved from tmp path")
	ErrFileInvalidTargetType  = errx.InvalidArg("FILE_INVALID_TARGET_TYPE", "invalid target type for file move")
	ErrFileUploadFailed       = errx.Internal("FILE_UPLOAD_FAILED", "file upload failed")
	ErrFileCreateFailed       = errx.Internal("FILE_CREATE_FAILED", "failed to create file")
	ErrFileUpdateFailed       = errx.Internal("FILE_UPDATE_FAILED", "failed to update file")
	ErrFileDeleteFailed       = errx.Internal("FILE_DELETE_FAILED", "failed to delete file")
	ErrFileQueryFailed        = errx.Internal("FILE_QUERY_FAILED", "file query failed")
	ErrInvalidFileID          = errx.InvalidArg("FILE_INVALID_ID", "invalid file id")
	ErrFileRequestBodyInvalid = errx.InvalidArg("FILE_REQUEST_BODY_INVALID", "invalid file request body")
)

/*
!FILE_NOT_FOUND
*en<File not found>
*zh<文件不存在>
*fr<fichier introuvable>

!FILE_FILE_PATH_REQUIRED
*en<File path is required>
*zh<文件路径必填>
*fr<chemin du fichier requis>

!FILE_FILE_NAME_REQUIRED
*en<File name is required>
*zh<文件名必填>
*fr<nom du fichier requis>

!FILE_FILE_SIZE_INVALID
*en<File size must be greater than 0>
*zh<文件大小必须大于 0>
*fr<taille du fichier doit être supérieure à 0>

!FILE_MIME_TYPE_REQUIRED
*en<File mime type is required>
*zh<文件 MIME 类型必填>
*fr<type MIME requis>

!FILE_UPLOADER_ID_REQUIRED
*en<File uploader id is required>
*zh<文件上传者 ID 必填>
*fr<ID uploader requis>

!FILE_ALREADY_EXISTS
*en<File already exists>
*zh<文件已存在>
*fr<fichier déjà existant>

!FILE_MOVE_ONLY_FROM_TMP
*en<File can only be moved from tmp path>
*zh<仅支持从临时路径移动文件>
*fr<le fichier ne peut être déplacé que depuis le chemin temporaire>

!FILE_INVALID_TARGET_TYPE
*en<Invalid target type for file move>
*zh<文件移动目标类型无效>
*fr<type de cible invalide pour le déplacement du fichier>

!FILE_UPLOAD_FAILED
*en<File upload failed>
*zh<文件上传失败>
*fr<échec du téléchargement du fichier>

!FILE_CREATE_FAILED
*en<Failed to create file>
*zh<创建文件记录失败>
*fr<échec de la création du fichier>

!FILE_UPDATE_FAILED
*en<Failed to update file>
*zh<更新文件失败>
*fr<échec de la mise à jour du fichier>

!FILE_DELETE_FAILED
*en<Failed to delete file>
*zh<删除文件失败>
*fr<échec de la suppression du fichier>

!FILE_QUERY_FAILED
*en<File query failed>
*zh<文件查询失败>
*fr<échec de la requête fichier>

!FILE_INVALID_ID
*en<Invalid file id>
*zh<文件 ID 无效>
*fr<ID de fichier invalide>

!FILE_REQUEST_BODY_INVALID
*en<Invalid file request body>
*zh<文件请求体无效>
*fr<corps de requête fichier invalide>
*/
