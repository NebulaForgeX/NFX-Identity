package access

import "nfxid/pkgs/errx"

const (
	CodeApplicationRoleNotFound = "APPLICATION_ROLE_NOT_FOUND"
	CodeApplicationIDRequired   = "APPLICATION_ID_REQUIRED"
	CodeRoleKeyRequired         = "ROLE_KEY_REQUIRED"
)

var (
	ErrApplicationRoleNotFound              = errx.NotFound(CodeApplicationRoleNotFound, "application role not found")
	ErrApplicationRoleApplicationIDRequired = errx.InvalidArg(CodeApplicationIDRequired, "application_id is required")
	ErrApplicationRoleRoleKeyRequired       = errx.InvalidArg(CodeRoleKeyRequired, "role_key is required")
)
