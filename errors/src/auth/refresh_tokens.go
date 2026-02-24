package auth

import "nfxid/pkgs/errx"

const (
	CodeRefreshTokenNotFound = "REFRESH_TOKEN_NOT_FOUND"
	CodeTokenIDRequired      = "TOKEN_ID_REQUIRED"
	CodeTokenIDAlreadyExists = "TOKEN_ID_ALREADY_EXISTS"
	CodeTokenAlreadyRevoked  = "TOKEN_ALREADY_REVOKED"
	CodeTokenExpired         = "TOKEN_EXPIRED"
)

var (
	ErrRefreshTokenNotFound = errx.NotFound(CodeRefreshTokenNotFound, "refresh token not found")
	ErrTokenIDRequired      = errx.InvalidArg(CodeTokenIDRequired, "token id is required")
	ErrTokenIDAlreadyExists = errx.Conflict(CodeTokenIDAlreadyExists, "token id already exists")
	ErrTokenAlreadyRevoked  = errx.Conflict(CodeTokenAlreadyRevoked, "token already revoked")
	ErrTokenExpired         = errx.Expired(CodeTokenExpired, "token expired")
)
