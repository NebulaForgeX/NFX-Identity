package trusted_devices

import (
	"time"
)

func (td *TrustedDevice) UpdateLastUsed() error {
	now := time.Now().UTC()
	td.state.LastUsedAt = now
	td.state.UpdatedAt = now
	return nil
}

func (td *TrustedDevice) UpdateTrustedUntil(trustedUntil time.Time) error {
	td.state.TrustedUntil = trustedUntil
	td.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (td *TrustedDevice) UpdateDeviceName(deviceName *string) error {
	td.state.DeviceName = deviceName
	td.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (td *TrustedDevice) IsTrusted() bool {
	return time.Now().UTC().Before(td.TrustedUntil())
}
