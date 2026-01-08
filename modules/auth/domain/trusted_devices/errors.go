package trusted_devices

import "errors"

var (
	ErrTrustedDeviceNotFound           = errors.New("trusted device not found")
	ErrDeviceIDRequired                = errors.New("device id is required")
	ErrUserIDRequired                  = errors.New("user id is required")
	ErrTenantIDRequired                = errors.New("tenant id is required")
	ErrDeviceFingerprintHashRequired   = errors.New("device fingerprint hash is required")
	ErrTrustedUntilRequired            = errors.New("trusted until is required")
	ErrTrustedDeviceAlreadyExists      = errors.New("trusted device already exists")
)
