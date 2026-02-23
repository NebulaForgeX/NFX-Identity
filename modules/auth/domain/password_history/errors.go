package password_history

import "nfxid/pkgs/errx"

var (
	ErrPasswordHistoryNotFound = errx.NotFound("PASSWORD_HISTORY_NOT_FOUND", "password history not found")
	ErrUserIDRequired          = errx.InvalidArg("USER_ID_REQUIRED", "user id is required")
	ErrTenantIDRequired        = errx.InvalidArg("TENANT_ID_REQUIRED", "tenant id is required")
	ErrPasswordHashRequired   = errx.InvalidArg("PASSWORD_HASH_REQUIRED", "password hash is required")
)
