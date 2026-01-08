package password_resets

import "errors"

var (
	ErrPasswordResetNotFound   = errors.New("password reset not found")
	ErrResetIDRequired         = errors.New("reset id is required")
	ErrUserIDRequired          = errors.New("user id is required")
	ErrTenantIDRequired        = errors.New("tenant id is required")
	ErrDeliveryRequired        = errors.New("delivery is required")
	ErrCodeHashRequired        = errors.New("code hash is required")
	ErrExpiresAtRequired       = errors.New("expires at is required")
	ErrResetIDAlreadyExists    = errors.New("reset id already exists")
	ErrInvalidResetDelivery    = errors.New("invalid reset delivery")
	ErrInvalidResetStatus      = errors.New("invalid reset status")
	ErrResetAlreadyUsed        = errors.New("reset already used")
	ErrResetExpired            = errors.New("reset expired")
)
