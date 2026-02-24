package auth

import "nfxid/pkgs/errx"

const (
	CodePasswordHistoryNotFound = "PASSWORD_HISTORY_NOT_FOUND"
)

var (
	ErrPasswordHistoryNotFound = errx.NotFound(CodePasswordHistoryNotFound, "password history not found")
)

/*
!PASSWORD_HISTORY_NOT_FOUND
*en<password history not found>
*zh<密码历史不存在>
*fr<historique du mot de passe introuvable>

*/
