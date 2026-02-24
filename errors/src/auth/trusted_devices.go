package auth

import "nfxid/pkgs/errx"

const (
	CodeTrustedDeviceNotFound         = "TRUSTED_DEVICE_NOT_FOUND"
	CodeDeviceIDRequired              = "DEVICE_ID_REQUIRED"
	CodeDeviceFingerprintHashRequired = "DEVICE_FINGERPRINT_HASH_REQUIRED"
	CodeTrustedUntilRequired          = "TRUSTED_UNTIL_REQUIRED"
	CodeTrustedDeviceAlreadyExists    = "TRUSTED_DEVICE_ALREADY_EXISTS"
)

var (
	ErrTrustedDeviceNotFound         = errx.NotFound(CodeTrustedDeviceNotFound, "trusted device not found")
	ErrDeviceIDRequired              = errx.InvalidArg(CodeDeviceIDRequired, "device id is required")
	ErrDeviceFingerprintHashRequired = errx.InvalidArg(CodeDeviceFingerprintHashRequired, "device fingerprint hash is required")
	ErrTrustedUntilRequired          = errx.InvalidArg(CodeTrustedUntilRequired, "trusted until is required")
	ErrTrustedDeviceAlreadyExists    = errx.Conflict(CodeTrustedDeviceAlreadyExists, "trusted device already exists")
)
