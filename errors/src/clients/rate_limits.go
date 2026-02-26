package clients

import "nfxidentity/pkgs/errx"

const (
	CodeRateLimitNotFound      = "RATE_LIMIT_NOT_FOUND"
	CodeLimitTypeRequired      = "LIMIT_TYPE_REQUIRED"
	CodeLimitValueRequired     = "LIMIT_VALUE_REQUIRED"
	CodeWindowSecondsRequired  = "WINDOW_SECONDS_REQUIRED"
	CodeRateLimitAlreadyExists = "RATE_LIMIT_ALREADY_EXISTS"
	CodeInvalidRateLimitType   = "INVALID_RATE_LIMIT_TYPE"
	CodeInvalidStatus          = "INVALID_STATUS"
)

var (
	ErrRateLimitNotFound      = errx.NotFound(CodeRateLimitNotFound, "rate limit not found")
	ErrLimitTypeRequired      = errx.InvalidArg(CodeLimitTypeRequired, "limit type is required")
	ErrLimitValueRequired     = errx.InvalidArg(CodeLimitValueRequired, "limit value is required")
	ErrWindowSecondsRequired  = errx.InvalidArg(CodeWindowSecondsRequired, "window seconds is required")
	ErrRateLimitAlreadyExists = errx.Conflict(CodeRateLimitAlreadyExists, "rate limit already exists")
	ErrInvalidRateLimitType   = errx.InvalidArg(CodeInvalidRateLimitType, "invalid rate limit type")
	ErrInvalidStatus          = errx.InvalidArg(CodeInvalidStatus, "invalid status")
)

/*
!RATE_LIMIT_NOT_FOUND
*en<rate limit not found>
*zh<速率限制不存在>
*fr<limite de débit introuvable>

!LIMIT_TYPE_REQUIRED
*en<limit type required>
*zh<限制类型为必填>
*fr<type de limite requis>

!LIMIT_VALUE_REQUIRED
*en<limit value required>
*zh<限制值为必填>
*fr<valeur de limite requise>

!WINDOW_SECONDS_REQUIRED
*en<window seconds required>
*zh<窗口秒数为必填>
*fr<fenêtre en secondes requise>

!RATE_LIMIT_ALREADY_EXISTS
*en<rate limit already exists>
*zh<速率限制已存在>
*fr<limite de débit existe déjà>

!INVALID_RATE_LIMIT_TYPE
*en<invalid rate limit type>
*zh<无效的速率限制类型>
*fr<type de limite de débit invalide>

!INVALID_STATUS
*en<invalid status>
*zh<无效的状态>
*fr<statut invalide>

*/
