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

/*
!APPLICATION_ROLE_NOT_FOUND
*en<application role not found>
*zh<应用角色不存在>
*fr<application role non trouvé>

!APPLICATION_ID_REQUIRED
*en<application id is required>
*zh<应用ID是必填的>
*fr<application id est requis>

!ROLE_KEY_REQUIRED
*en<role key is required>
*zh<角色键是必填的>
*fr<role key est requis>
*/