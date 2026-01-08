package trusted_devices

import "github.com/google/uuid"

func (td *TrustedDevice) Validate() error {
	if td.DeviceID() == "" {
		return ErrDeviceIDRequired
	}
	if td.UserID() == uuid.Nil {
		return ErrUserIDRequired
	}
	if td.TenantID() == uuid.Nil {
		return ErrTenantIDRequired
	}
	if td.DeviceFingerprintHash() == "" {
		return ErrDeviceFingerprintHashRequired
	}
	if td.TrustedUntil().IsZero() {
		return ErrTrustedUntilRequired
	}
	return nil
}
