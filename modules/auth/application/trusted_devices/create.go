package trusted_devices

import (
	"context"
	"time"
	trustedDeviceCommands "nfxid/modules/auth/application/trusted_devices/commands"
	trustedDeviceDomain "nfxid/modules/auth/domain/trusted_devices"

	"github.com/google/uuid"
)

// CreateTrustedDevice 创建受信任设备
func (s *Service) CreateTrustedDevice(ctx context.Context, cmd trustedDeviceCommands.CreateTrustedDeviceCmd) (uuid.UUID, error) {
	// Parse trusted until
	trustedUntil, err := time.Parse(time.RFC3339, cmd.TrustedUntil)
	if err != nil {
		return uuid.Nil, err
	}

	// Create domain entity
	trustedDevice, err := trustedDeviceDomain.NewTrustedDevice(trustedDeviceDomain.NewTrustedDeviceParams{
		DeviceID:             cmd.DeviceID,
		UserID:               cmd.UserID,
		TenantID:             cmd.TenantID,
		DeviceFingerprintHash: cmd.DeviceFingerprintHash,
		DeviceName:           cmd.DeviceName,
		TrustedUntil:         trustedUntil,
		IP:                   cmd.IP,
		UAHash:               cmd.UAHash,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.trustedDeviceRepo.Create.New(ctx, trustedDevice); err != nil {
		return uuid.Nil, err
	}

	return trustedDevice.ID(), nil
}
