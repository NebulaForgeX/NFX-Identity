package password_resets

import "nfxid/pkgs/errx"

var (
	ErrPasswordResetNotFound  = errx.NotFound("PASSWORD_RESET_NOT_FOUND", "password reset not found")
	ErrResetIDRequired       = errx.InvalidArg("RESET_ID_REQUIRED", "reset id is required")
	ErrUserIDRequired        = errx.InvalidArg("USER_ID_REQUIRED", "user id is required")
	ErrTenantIDRequired      = errx.InvalidArg("TENANT_ID_REQUIRED", "tenant id is required")
	ErrDeliveryRequired      = errx.InvalidArg("DELIVERY_REQUIRED", "delivery is required")
	ErrCodeHashRequired      = errx.InvalidArg("CODE_HASH_REQUIRED", "code hash is required")
	ErrExpiresAtRequired     = errx.InvalidArg("EXPIRES_AT_REQUIRED", "expires at is required")
	ErrResetIDAlreadyExists  = errx.Conflict("RESET_ID_ALREADY_EXISTS", "reset id already exists")
	ErrInvalidResetDelivery  = errx.InvalidArg("INVALID_RESET_DELIVERY", "invalid reset delivery")
	ErrInvalidResetStatus    = errx.InvalidArg("INVALID_RESET_STATUS", "invalid reset status")
	ErrResetAlreadyUsed     = errx.Conflict("RESET_ALREADY_USED", "reset already used")
	ErrResetExpired         = errx.Expired("RESET_EXPIRED", "reset expired")
)
