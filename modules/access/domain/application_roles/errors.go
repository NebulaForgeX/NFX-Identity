package application_roles

import "nfxid/pkgs/errx"

var (
	ErrApplicationRoleNotFound             = errx.NotFound("APPLICATION_ROLE_NOT_FOUND", "application role not found")
	ErrApplicationRoleApplicationIDRequired = errx.InvalidArg("APPLICATION_ID_REQUIRED", "application_id is required")
	ErrApplicationRoleRoleKeyRequired      = errx.InvalidArg("ROLE_KEY_REQUIRED", "role_key is required")
)
