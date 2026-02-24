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

/*
!REFRESH_TOKEN_NOT_FOUND
*en<refresh token not found>
*zh<刷新令牌不存在>
*fr<jeton de rafraîchissement introuvable>

!TOKEN_ID_REQUIRED
*en<token id required>
*zh<令牌 ID 为必填>
*fr<id du jeton requis>

!TOKEN_ID_ALREADY_EXISTS
*en<token id already exists>
*zh<令牌 ID 已存在>
*fr<id du jeton existe déjà>

!TOKEN_ALREADY_REVOKED
*en<token already revoked>
*zh<令牌已撤销>
*fr<jeton déjà révoqué>

!TOKEN_EXPIRED
*en<token expired>
*zh<令牌已过期>
*fr<jeton expiré>

*/
