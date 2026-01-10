package trusted_devices

import (
	trustedDeviceDomain "nfxid/modules/auth/domain/trusted_devices"
)

type Service struct {
	trustedDeviceRepo *trustedDeviceDomain.Repo
}

func NewService(
	trustedDeviceRepo *trustedDeviceDomain.Repo,
) *Service {
	return &Service{
		trustedDeviceRepo: trustedDeviceRepo,
	}
}
