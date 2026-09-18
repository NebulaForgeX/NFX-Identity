package auth

import "nfxidentity/pkgs/errx"

var (
	ErrRefreshTokenNotFound          = errx.NotFound("REFRESH_TOKEN_NOT_FOUND", "refresh token not found")
	ErrRefreshTokenAccountIDInvalid  = errx.InvalidArg("REFRESH_TOKEN_ACCOUNT_ID_INVALID", "invalid refresh token account id")
	ErrRefreshTokenTokenHashRequired = errx.InvalidArg("REFRESH_TOKEN_HASH_REQUIRED", "refresh token hash is required")
	ErrRefreshTokenExpiresAtRequired = errx.InvalidArg("REFRESH_TOKEN_EXPIRES_AT_REQUIRED", "refresh token expires at is required")
	ErrRefreshTokenExpiresAtInvalid  = errx.InvalidArg("REFRESH_TOKEN_EXPIRES_AT_INVALID", "invalid refresh token expires at")
	ErrRefreshTokenAlreadyRevoked    = errx.InvalidArg("REFRESH_TOKEN_ALREADY_REVOKED", "refresh token already revoked")
)

/*
!REFRESH_TOKEN_NOT_FOUND
*en<Refresh token not found>
*zh<刷新令牌不存在>
*fr<Jeton de rafraîchissement introuvable>

!REFRESH_TOKEN_ACCOUNT_ID_INVALID
*en<Invalid refresh token account id>
*zh<刷新令牌账号 ID 无效>
*fr<Identifiant de compte du jeton de rafraîchissement invalide>

!REFRESH_TOKEN_HASH_REQUIRED
*en<Refresh token hash is required>
*zh<刷新令牌哈希必填>
*fr<Hash du jeton de rafraîchissement requis>

!REFRESH_TOKEN_EXPIRES_AT_REQUIRED
*en<Refresh token expires at is required>
*zh<刷新令牌过期时间必填>
*fr<Date d'expiration du jeton requise>

!REFRESH_TOKEN_EXPIRES_AT_INVALID
*en<Invalid refresh token expires at>
*zh<刷新令牌过期时间无效>
*fr<Date d'expiration du jeton invalide>

!REFRESH_TOKEN_ALREADY_REVOKED
*en<Refresh token already revoked>
*zh<刷新令牌已撤销>
*fr<Jeton de rafraîchissement déjà révoqué>
*/
