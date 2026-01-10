package results

import (
	"time"

	"nfxid/modules/auth/domain/trusted_devices"

	"github.com/google/uuid"
)

type TrustedDeviceRO struct {
	ID                    uuid.UUID
	DeviceID              string
	UserID                uuid.UUID
	TenantID              uuid.UUID
	DeviceFingerprintHash string
	DeviceName            *string
	TrustedUntil          time.Time
	LastUsedAt            time.Time
	IP                    *string
	UAHash                *string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

// TrustedDeviceMapper 将 Domain TrustedDevice 转换为 Application TrustedDeviceRO
func TrustedDeviceMapper(td *trusted_devices.TrustedDevice) TrustedDeviceRO {
	if td == nil {
		return TrustedDeviceRO{}
	}

	return TrustedDeviceRO{
		ID:                    td.ID(),
		DeviceID:              td.DeviceID(),
		UserID:                td.UserID(),
		TenantID:              td.TenantID(),
		DeviceFingerprintHash: td.DeviceFingerprintHash(),
		DeviceName:            td.DeviceName(),
		TrustedUntil:          td.TrustedUntil(),
		LastUsedAt:            td.LastUsedAt(),
		IP:                    td.IP(),
		UAHash:                td.UAHash(),
		CreatedAt:             td.CreatedAt(),
		UpdatedAt:             td.UpdatedAt(),
	}
}
