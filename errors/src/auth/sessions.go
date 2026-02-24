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
