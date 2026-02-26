package trusted_devices

import (
	trustedDeviceDomain "nfxidentity/modules/auth/domain/trusted_devices"
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
