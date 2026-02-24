package auth

import "nfxid/pkgs/errx"

const (
	CodePasswordHistoryNotFound = "PASSWORD_HISTORY_NOT_FOUND"
)

var (
	ErrPasswordHistoryNotFound = errx.NotFound(CodePasswordHistoryNotFound, "password history not found")
)
