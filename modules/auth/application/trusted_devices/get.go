package trusted_devices

import (
	"context"
	trustedDeviceResult "nfxid/modules/auth/application/trusted_devices/results"

	"github.com/google/uuid"
)

// GetTrustedDevice 根据ID获取受信任设备
func (s *Service) GetTrustedDevice(ctx context.Context, trustedDeviceID uuid.UUID) (trustedDeviceResult.TrustedDeviceRO, error) {
	domainEntity, err := s.trustedDeviceRepo.Get.ByID(ctx, trustedDeviceID)
	if err != nil {
		return trustedDeviceResult.TrustedDeviceRO{}, err
	}
	return trustedDeviceResult.TrustedDeviceMapper(domainEntity), nil
}

// GetTrustedDeviceByDeviceID 根据DeviceID获取受信任设备
func (s *Service) GetTrustedDeviceByDeviceID(ctx context.Context, deviceID string) (trustedDeviceResult.TrustedDeviceRO, error) {
	domainEntity, err := s.trustedDeviceRepo.Get.ByDeviceID(ctx, deviceID)
	if err != nil {
		return trustedDeviceResult.TrustedDeviceRO{}, err
	}
	return trustedDeviceResult.TrustedDeviceMapper(domainEntity), nil
}
