package directory

import "nfxid/pkgs/errx"

const (
	CodeUserEmailNotFound  = "USER_EMAIL_NOT_FOUND"
	CodeEmailRequired      = "EMAIL_REQUIRED"
	CodeEmailAlreadyExists = "EMAIL_ALREADY_EXISTS"
	CodeInvalidEmail       = "INVALID_EMAIL"
)

var (
	ErrUserEmailNotFound  = errx.NotFound(CodeUserEmailNotFound, "user email not found")
	ErrEmailRequired      = errx.InvalidArg(CodeEmailRequired, "email is required")
	ErrEmailAlreadyExists = errx.Conflict(CodeEmailAlreadyExists, "email already exists")
	ErrInvalidEmail       = errx.InvalidArg(CodeInvalidEmail, "invalid email format")
)

/*
!USER_EMAIL_NOT_FOUND
*en<user email not found>
*zh<用户邮箱不存在>
*fr<email utilisateur introuvable>

!EMAIL_REQUIRED
*en<email required>
*zh<邮箱为必填>
*fr<email requis>

!EMAIL_ALREADY_EXISTS
*en<email already exists>
*zh<邮箱已存在>
*fr<email existe déjà>

!INVALID_EMAIL
*en<invalid email>
*zh<无效的邮箱格式>
*fr<email invalide>

*/
