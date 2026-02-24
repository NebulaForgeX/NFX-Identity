package auth

import "nfxid/pkgs/errx"

const (
	CodeLoginAttemptNotFound = "LOGIN_ATTEMPT_NOT_FOUND"
	CodeIdentifierRequired   = "IDENTIFIER_REQUIRED"
	CodeInvalidFailureCode   = "INVALID_FAILURE_CODE"
)

var (
	ErrLoginAttemptNotFound = errx.NotFound(CodeLoginAttemptNotFound, "login attempt not found")
	ErrIdentifierRequired   = errx.InvalidArg(CodeIdentifierRequired, "identifier is required")
	ErrInvalidFailureCode   = errx.InvalidArg(CodeInvalidFailureCode, "invalid failure code")
)

/*
!LOGIN_ATTEMPT_NOT_FOUND
*en<login attempt not found>
*zh<登录尝试不存在>
*fr<tentative de connexion introuvable>

!IDENTIFIER_REQUIRED
*en<identifier required>
*zh<标识符为必填>
*fr<identifiant requis>

!INVALID_FAILURE_CODE
*en<invalid failure code>
*zh<无效的失败代码>
*fr<code d'échec invalide>

*/
