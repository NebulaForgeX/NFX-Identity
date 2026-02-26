package auth

import "nfxidentity/pkgs/errx"

const (
	CodeMFAFactorNotFound     = "MFA_FACTOR_NOT_FOUND"
	CodeFactorIDRequired      = "FACTOR_ID_REQUIRED"
	CodeTypeRequired          = "TYPE_REQUIRED"
	CodeFactorIDAlreadyExists = "FACTOR_ID_ALREADY_EXISTS"
	CodeInvalidMFAType        = "INVALID_MFA_TYPE"
)

var (
	ErrMFAFactorNotFound     = errx.NotFound(CodeMFAFactorNotFound, "mfa factor not found")
	ErrFactorIDRequired      = errx.InvalidArg(CodeFactorIDRequired, "factor id is required")
	ErrTypeRequired          = errx.InvalidArg(CodeTypeRequired, "type is required")
	ErrFactorIDAlreadyExists = errx.Conflict(CodeFactorIDAlreadyExists, "factor id already exists")
	ErrInvalidMFAType        = errx.InvalidArg(CodeInvalidMFAType, "invalid mfa type")
)

/*
!MFA_FACTOR_NOT_FOUND
*en<mfa factor not found>
*zh<MFA 因子不存在>
*fr<facteur MFA introuvable>

!FACTOR_ID_REQUIRED
*en<factor id required>
*zh<因子 ID 为必填>
*fr<id du facteur requis>

!TYPE_REQUIRED
*en<type required>
*zh<类型为必填>
*fr<type requis>

!FACTOR_ID_ALREADY_EXISTS
*en<factor id already exists>
*zh<因子 ID 已存在>
*fr<id du facteur existe déjà>

!INVALID_MFA_TYPE
*en<invalid mfa type>
*zh<无效的 MFA 类型>
*fr<type MFA invalide>

*/
