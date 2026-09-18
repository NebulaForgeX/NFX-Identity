package asset

import "nfxidentity/pkgs/errx"

// 音频（asset."Audios"）相关
var (
	ErrAudioNotFound           = errx.NotFound("AUDIO_NOT_FOUND", "audio not found")
	ErrAudioFilePathRequired   = errx.InvalidArg("AUDIO_FILE_PATH_REQUIRED", "audio file path is required")
	ErrAudioFileNameRequired   = errx.InvalidArg("AUDIO_FILE_NAME_REQUIRED", "audio file name is required")
	ErrAudioFileSizeInvalid    = errx.InvalidArg("AUDIO_FILE_SIZE_INVALID", "audio file size must be greater than 0")
	ErrAudioMimeTypeRequired   = errx.InvalidArg("AUDIO_MIME_TYPE_REQUIRED", "audio mime type is required")
	ErrAudioUploaderIDRequired = errx.InvalidArg("AUDIO_UPLOADER_ID_REQUIRED", "audio uploader id is required")
	ErrAudioAlreadyExists      = errx.Conflict("AUDIO_ALREADY_EXISTS", "audio already exists")
	ErrAudioMoveOnlyFromTmp    = errx.InvalidArg("AUDIO_MOVE_ONLY_FROM_TMP", "audio can only be moved from tmp path")
	ErrAudioInvalidTargetType  = errx.InvalidArg("AUDIO_INVALID_TARGET_TYPE", "invalid target type for audio move")
	ErrAudioUploadFailed       = errx.Internal("AUDIO_UPLOAD_FAILED", "audio upload failed")
	ErrAudioCreateFailed       = errx.Internal("AUDIO_CREATE_FAILED", "failed to create audio")
	ErrAudioUpdateFailed       = errx.Internal("AUDIO_UPDATE_FAILED", "failed to update audio")
	ErrAudioDeleteFailed       = errx.Internal("AUDIO_DELETE_FAILED", "failed to delete audio")
	ErrAudioQueryFailed        = errx.Internal("AUDIO_QUERY_FAILED", "audio query failed")
	ErrInvalidAudioID          = errx.InvalidArg("AUDIO_INVALID_ID", "invalid audio id")
	ErrAudioRequestBodyInvalid = errx.InvalidArg("AUDIO_REQUEST_BODY_INVALID", "invalid audio request body")
)

/*
!AUDIO_NOT_FOUND
*en<Audio not found>
*zh<音频不存在>
*fr<audio introuvable>

!AUDIO_FILE_PATH_REQUIRED
*en<Audio file path is required>
*zh<音频文件路径必填>
*fr<chemin de l'audio requis>

!AUDIO_FILE_NAME_REQUIRED
*en<Audio file name is required>
*zh<音频文件名必填>
*fr<nom du fichier audio requis>

!AUDIO_FILE_SIZE_INVALID
*en<Audio file size must be greater than 0>
*zh<音频文件大小必须大于 0>
*fr<taille de l'audio doit être supérieure à 0>

!AUDIO_MIME_TYPE_REQUIRED
*en<Audio mime type is required>
*zh<音频 MIME 类型必填>
*fr<type MIME de l'audio requis>

!AUDIO_UPLOADER_ID_REQUIRED
*en<Audio uploader id is required>
*zh<音频上传者 ID 必填>
*fr<ID uploader requis>

!AUDIO_ALREADY_EXISTS
*en<Audio already exists>
*zh<音频已存在>
*fr<audio déjà existant>

!AUDIO_MOVE_ONLY_FROM_TMP
*en<Audio can only be moved from tmp path>
*zh<仅支持从临时路径移动音频>
*fr<l'audio ne peut être déplacé que depuis le chemin temporaire>

!AUDIO_INVALID_TARGET_TYPE
*en<Invalid target type for audio move>
*zh<音频移动目标类型无效>
*fr<type de cible invalide pour le déplacement de l'audio>

!AUDIO_UPLOAD_FAILED
*en<Audio upload failed>
*zh<音频上传失败>
*fr<échec du téléchargement de l'audio>

!AUDIO_CREATE_FAILED
*en<Failed to create audio>
*zh<创建音频记录失败>
*fr<échec de la création de l'audio>

!AUDIO_UPDATE_FAILED
*en<Failed to update audio>
*zh<更新音频失败>
*fr<échec de la mise à jour de l'audio>

!AUDIO_DELETE_FAILED
*en<Failed to delete audio>
*zh<删除音频失败>
*fr<échec de la suppression de l'audio>

!AUDIO_QUERY_FAILED
*en<Audio query failed>
*zh<音频查询失败>
*fr<échec de la requête audio>

!AUDIO_INVALID_ID
*en<Invalid audio id>
*zh<音频 ID 无效>
*fr<ID audio invalide>

!AUDIO_REQUEST_BODY_INVALID
*en<Invalid audio request body>
*zh<音频请求体无效>
*fr<corps de requête audio invalide>
*/
