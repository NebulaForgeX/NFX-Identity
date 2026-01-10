package trusted_devices

import (
	"context"
	trustedDeviceCommands "nfxid/modules/auth/application/trusted_devices/commands"
)

// DeleteTrustedDevice 删除受信任设备
func (s *Service) DeleteTrustedDevice(ctx context.Context, cmd trustedDeviceCommands.DeleteTrustedDeviceCmd) error {
	// Delete from repository (hard delete)
	return s.trustedDeviceRepo.Delete.ByDeviceID(ctx, cmd.DeviceID)
}
