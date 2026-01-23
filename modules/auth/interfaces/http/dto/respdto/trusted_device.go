package respdto

import (
	"time"

	trustedDeviceAppResult "nfxid/modules/auth/application/trusted_devices/results"

	"github.com/google/uuid"
)

type TrustedDeviceDTO struct {
	ID                    uuid.UUID  `json:"id"`
	DeviceID              string     `json:"device_id"`
	UserID                uuid.UUID  `json:"user_id"`
	TenantID              uuid.UUID  `json:"tenant_id"`
	DeviceFingerprintHash string     `json:"device_fingerprint_hash"`
	DeviceName            *string    `json:"device_name,omitempty"`
	TrustedUntil          time.Time  `json:"trusted_until"`
	LastUsedAt            time.Time  `json:"last_used_at"`
	IP                    *string    `json:"ip,omitempty"`
	UAHash                *string    `json:"ua_hash,omitempty"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
}

// TrustedDeviceROToDTO converts application TrustedDeviceRO to response DTO
func TrustedDeviceROToDTO(v *trustedDeviceAppResult.TrustedDeviceRO) *TrustedDeviceDTO {
	if v == nil {
		return nil
	}

	return &TrustedDeviceDTO{
		ID:                    v.ID,
		DeviceID:              v.DeviceID,
		UserID:                v.UserID,
		TenantID:              v.TenantID,
		DeviceFingerprintHash: v.DeviceFingerprintHash,
		DeviceName:            v.DeviceName,
		TrustedUntil:          v.TrustedUntil,
		LastUsedAt:            v.LastUsedAt,
		IP:                    v.IP,
		UAHash:                v.UAHash,
		CreatedAt:             v.CreatedAt,
		UpdatedAt:             v.UpdatedAt,
	}
}
