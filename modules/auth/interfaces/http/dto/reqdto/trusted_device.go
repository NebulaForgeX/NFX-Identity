package reqdto

import (
	trustedDeviceAppCommands "nfxid/modules/auth/application/trusted_devices/commands"

	"github.com/google/uuid"
)

type TrustedDeviceCreateRequestDTO struct {
	DeviceID              string    `json:"device_id" validate:"required"`
	UserID                uuid.UUID `json:"user_id" validate:"required"`
	TenantID              uuid.UUID `json:"tenant_id" validate:"required"`
	DeviceFingerprintHash string    `json:"device_fingerprint_hash" validate:"required"`
	DeviceName            *string   `json:"device_name,omitempty"`
	TrustedUntil          string    `json:"trusted_until" validate:"required"`
	IP                    *string   `json:"ip,omitempty"`
	UAHash                *string   `json:"ua_hash,omitempty"`
}

type TrustedDeviceByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

func (r *TrustedDeviceCreateRequestDTO) ToCreateCmd() trustedDeviceAppCommands.CreateTrustedDeviceCmd {
	return trustedDeviceAppCommands.CreateTrustedDeviceCmd{
		DeviceID:              r.DeviceID,
		UserID:                r.UserID,
		TenantID:              r.TenantID,
		DeviceFingerprintHash: r.DeviceFingerprintHash,
		DeviceName:            r.DeviceName,
		TrustedUntil:          r.TrustedUntil,
		IP:                    r.IP,
		UAHash:                r.UAHash,
	}
}
