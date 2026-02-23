package trusted_devices

import "nfxid/pkgs/errx"

var (
	ErrTrustedDeviceNotFound         = errx.NotFound("TRUSTED_DEVICE_NOT_FOUND", "trusted device not found")
	ErrDeviceIDRequired              = errx.InvalidArg("DEVICE_ID_REQUIRED", "device id is required")
	ErrUserIDRequired                = errx.InvalidArg("USER_ID_REQUIRED", "user id is required")
	ErrTenantIDRequired              = errx.InvalidArg("TENANT_ID_REQUIRED", "tenant id is required")
	ErrDeviceFingerprintHashRequired = errx.InvalidArg("DEVICE_FINGERPRINT_HASH_REQUIRED", "device fingerprint hash is required")
	ErrTrustedUntilRequired          = errx.InvalidArg("TRUSTED_UNTIL_REQUIRED", "trusted until is required")
	ErrTrustedDeviceAlreadyExists    = errx.Conflict("TRUSTED_DEVICE_ALREADY_EXISTS", "trusted device already exists")
)
