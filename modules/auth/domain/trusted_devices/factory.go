package trusted_devices

import (
	"time"

	"github.com/google/uuid"
)

type NewTrustedDeviceParams struct {
	DeviceID             string
	UserID               uuid.UUID
	TenantID             uuid.UUID
	DeviceFingerprintHash string
	DeviceName           *string
	TrustedUntil         time.Time
	IP                   *string
	UAHash               *string
}

func NewTrustedDevice(p NewTrustedDeviceParams) (*TrustedDevice, error) {
	if err := validateTrustedDeviceParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewTrustedDeviceFromState(TrustedDeviceState{
		ID:                   id,
		DeviceID:             p.DeviceID,
		UserID:               p.UserID,
		TenantID:             p.TenantID,
		DeviceFingerprintHash: p.DeviceFingerprintHash,
		DeviceName:           p.DeviceName,
		TrustedUntil:         p.TrustedUntil,
		LastUsedAt:           now,
		IP:                   p.IP,
		UAHash:               p.UAHash,
		CreatedAt:            now,
		UpdatedAt:            now,
	}), nil
}

func NewTrustedDeviceFromState(st TrustedDeviceState) *TrustedDevice {
	return &TrustedDevice{state: st}
}

func validateTrustedDeviceParams(p NewTrustedDeviceParams) error {
	if p.DeviceID == "" {
		return ErrDeviceIDRequired
	}
	if p.UserID == uuid.Nil {
		return ErrUserIDRequired
	}
	if p.TenantID == uuid.Nil {
		return ErrTenantIDRequired
	}
	if p.DeviceFingerprintHash == "" {
		return ErrDeviceFingerprintHashRequired
	}
	if p.TrustedUntil.IsZero() {
		return ErrTrustedUntilRequired
	}
	return nil
}
