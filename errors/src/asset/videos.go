package asset

import "nfxidentity/pkgs/errx"

// 视频（asset."Videos"）相关
var (
	ErrVideoNotFound           = errx.NotFound("VIDEO_NOT_FOUND", "video not found")
	ErrVideoFilePathRequired   = errx.InvalidArg("VIDEO_FILE_PATH_REQUIRED", "video file path is required")
	ErrVideoFileNameRequired   = errx.InvalidArg("VIDEO_FILE_NAME_REQUIRED", "video file name is required")
	ErrVideoFileSizeInvalid    = errx.InvalidArg("VIDEO_FILE_SIZE_INVALID", "video file size must be greater than 0")
	ErrVideoMimeTypeRequired   = errx.InvalidArg("VIDEO_MIME_TYPE_REQUIRED", "video mime type is required")
	ErrVideoUploaderIDRequired = errx.InvalidArg("VIDEO_UPLOADER_ID_REQUIRED", "video uploader id is required")
	ErrVideoAlreadyExists      = errx.Conflict("VIDEO_ALREADY_EXISTS", "video already exists")
	ErrVideoMoveOnlyFromTmp    = errx.InvalidArg("VIDEO_MOVE_ONLY_FROM_TMP", "video can only be moved from tmp path")
	ErrVideoInvalidTargetType  = errx.InvalidArg("VIDEO_INVALID_TARGET_TYPE", "invalid target type for video move")
	ErrVideoUploadFailed       = errx.Internal("VIDEO_UPLOAD_FAILED", "video upload failed")
	ErrVideoCreateFailed       = errx.Internal("VIDEO_CREATE_FAILED", "failed to create video")
	ErrVideoUpdateFailed       = errx.Internal("VIDEO_UPDATE_FAILED", "failed to update video")
	ErrVideoDeleteFailed       = errx.Internal("VIDEO_DELETE_FAILED", "failed to delete video")
	ErrVideoQueryFailed        = errx.Internal("VIDEO_QUERY_FAILED", "video query failed")
	ErrInvalidVideoID          = errx.InvalidArg("VIDEO_INVALID_ID", "invalid video id")
	ErrVideoRequestBodyInvalid = errx.InvalidArg("VIDEO_REQUEST_BODY_INVALID", "invalid video request body")
)

/*
!VIDEO_NOT_FOUND
*en<Video not found>
*zh<视频不存在>
*fr<vidéo introuvable>

!VIDEO_FILE_PATH_REQUIRED
*en<Video file path is required>
*zh<视频文件路径必填>
*fr<chemin de la vidéo requis>

!VIDEO_FILE_NAME_REQUIRED
*en<Video file name is required>
*zh<视频文件名必填>
*fr<nom du fichier vidéo requis>

!VIDEO_FILE_SIZE_INVALID
*en<Video file size must be greater than 0>
*zh<视频文件大小必须大于 0>
*fr<taille de la vidéo doit être supérieure à 0>

!VIDEO_MIME_TYPE_REQUIRED
*en<Video mime type is required>
*zh<视频 MIME 类型必填>
*fr<type MIME de la vidéo requis>

!VIDEO_UPLOADER_ID_REQUIRED
*en<Video uploader id is required>
*zh<视频上传者 ID 必填>
*fr<ID uploader requis>

!VIDEO_ALREADY_EXISTS
*en<Video already exists>
*zh<视频已存在>
*fr<vidéo déjà existante>

!VIDEO_MOVE_ONLY_FROM_TMP
*en<Video can only be moved from tmp path>
*zh<仅支持从临时路径移动视频>
*fr<la vidéo ne peut être déplacée que depuis le chemin temporaire>

!VIDEO_INVALID_TARGET_TYPE
*en<Invalid target type for video move>
*zh<视频移动目标类型无效>
*fr<type de cible invalide pour le déplacement de la vidéo>

!VIDEO_UPLOAD_FAILED
*en<Video upload failed>
*zh<视频上传失败>
*fr<échec du téléchargement de la vidéo>

!VIDEO_CREATE_FAILED
*en<Failed to create video>
*zh<创建视频记录失败>
*fr<échec de la création de la vidéo>

!VIDEO_UPDATE_FAILED
*en<Failed to update video>
*zh<更新视频失败>
*fr<échec de la mise à jour de la vidéo>

!VIDEO_DELETE_FAILED
*en<Failed to delete video>
*zh<删除视频失败>
*fr<échec de la suppression de la vidéo>

!VIDEO_QUERY_FAILED
*en<Video query failed>
*zh<视频查询失败>
*fr<échec de la requête vidéo>

!VIDEO_INVALID_ID
*en<Invalid video id>
*zh<视频 ID 无效>
*fr<ID vidéo invalide>

!VIDEO_REQUEST_BODY_INVALID
*en<Invalid video request body>
*zh<视频请求体无效>
*fr<corps de requête vidéo invalide>
*/
