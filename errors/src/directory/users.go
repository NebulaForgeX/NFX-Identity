package directory

import "nfxid/pkgs/errx"

const (
	CodeUserNotFound          = "USER_NOT_FOUND"
	CodeUsernameRequired      = "USERNAME_REQUIRED"
	CodeUsernameAlreadyExists = "USERNAME_ALREADY_EXISTS"
	CodeInvalidUserStatus     = "INVALID_USER_STATUS"
)

var (
	ErrUserNotFound          = errx.NotFound(CodeUserNotFound, "user not found")
	ErrUsernameRequired      = errx.InvalidArg(CodeUsernameRequired, "username is required")
	ErrUsernameAlreadyExists = errx.Conflict(CodeUsernameAlreadyExists, "username already exists")
	ErrInvalidUserStatus     = errx.InvalidArg(CodeInvalidUserStatus, "invalid user status")
)

/*
!USER_NOT_FOUND
*en<user not found>
*zh<用户不存在>
*fr<utilisateur introuvable>

!USERNAME_REQUIRED
*en<username required>
*zh<用户名为必填>
*fr<nom d'utilisateur requis>

!USERNAME_ALREADY_EXISTS
*en<username already exists>
*zh<用户名已存在>
*fr<nom d'utilisateur existe déjà>

!INVALID_USER_STATUS
*en<invalid user status>
*zh<无效的用户状态>
*fr<statut utilisateur invalide>

*/
