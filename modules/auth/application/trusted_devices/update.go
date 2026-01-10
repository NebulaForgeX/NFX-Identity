package trusted_devices

import (
	"context"
	"time"
	trustedDeviceCommands "nfxid/modules/auth/application/trusted_devices/commands"
)

// UpdateLastUsed 更新最后使用时间
func (s *Service) UpdateLastUsed(ctx context.Context, cmd trustedDeviceCommands.UpdateLastUsedCmd) error {
	// Get domain entity
	trustedDevice, err := s.trustedDeviceRepo.Get.ByDeviceID(ctx, cmd.DeviceID)
	if err != nil {
		return err
	}

	// Update last used domain entity
	if err := trustedDevice.UpdateLastUsed(); err != nil {
		return err
	}

	// Save to repository
	return s.trustedDeviceRepo.Update.UpdateLastUsed(ctx, cmd.DeviceID)
}

// UpdateTrustedUntil 更新信任到期时间
func (s *Service) UpdateTrustedUntil(ctx context.Context, cmd trustedDeviceCommands.UpdateTrustedUntilCmd) error {
	// Get domain entity
	trustedDevice, err := s.trustedDeviceRepo.Get.ByDeviceID(ctx, cmd.DeviceID)
	if err != nil {
		return err
	}

	// Parse trusted until
	trustedUntil, err := time.Parse(time.RFC3339, cmd.TrustedUntil)
	if err != nil {
		return err
	}

	// Update trusted until domain entity
	if err := trustedDevice.UpdateTrustedUntil(trustedUntil); err != nil {
		return err
	}

	// Save to repository
	return s.trustedDeviceRepo.Update.UpdateTrustedUntil(ctx, cmd.DeviceID, trustedUntil)
}
