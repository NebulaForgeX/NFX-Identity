package trusted_devices

import (
	authErr "nfxid/errors/src/auth"

	"github.com/google/uuid"
)

func (td *TrustedDevice) Validate() error {
	if td.DeviceID() == "" {
		return authErr.ErrDeviceIDRequired
	}
	if td.UserID() == uuid.Nil {
		return authErr.ErrUserIDRequired
	}
	if td.TenantID() == uuid.Nil {
		return authErr.ErrTenantIDRequired
	}
	if td.DeviceFingerprintHash() == "" {
		return authErr.ErrDeviceFingerprintHashRequired
	}
	if td.TrustedUntil().IsZero() {
		return authErr.ErrTrustedUntilRequired
	}
	return nil
}
