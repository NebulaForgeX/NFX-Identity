package auth

import "nfxid/pkgs/errx"

const (
	CodeSessionNotFound        = "SESSION_NOT_FOUND"
	CodeSessionIDRequired      = "SESSION_ID_REQUIRED"
	CodeSessionIDAlreadyExists = "SESSION_ID_ALREADY_EXISTS"
	CodeSessionAlreadyRevoked  = "SESSION_ALREADY_REVOKED"
	CodeSessionExpired         = "SESSION_EXPIRED"
)

var (
	ErrSessionNotFound        = errx.NotFound(CodeSessionNotFound, "session not found")
	ErrSessionIDRequired      = errx.InvalidArg(CodeSessionIDRequired, "session id is required")
	ErrSessionIDAlreadyExists = errx.Conflict(CodeSessionIDAlreadyExists, "session id already exists")
	ErrSessionAlreadyRevoked  = errx.Conflict(CodeSessionAlreadyRevoked, "session already revoked")
	ErrSessionExpired         = errx.Expired(CodeSessionExpired, "session expired")
)

/*
!SESSION_NOT_FOUND
*en<session not found>
*zh<会话不存在>
*fr<session introuvable>

!SESSION_ID_REQUIRED
*en<session id required>
*zh<会话 ID 为必填>
*fr<id de session requis>

!SESSION_ID_ALREADY_EXISTS
*en<session id already exists>
*zh<会话 ID 已存在>
*fr<id de session existe déjà>

!SESSION_ALREADY_REVOKED
*en<session already revoked>
*zh<会话已撤销>
*fr<session déjà révoquée>

!SESSION_EXPIRED
*en<session expired>
*zh<会话已过期>
*fr<session expirée>

*/
