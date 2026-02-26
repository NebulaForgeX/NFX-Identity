package handler

import (
	"context"

	trustedDeviceApp "nfxidentity/modules/auth/application/trusted_devices"
	"nfxidentity/pkgs/errx"
	trusteddevicepb "nfxidentity/protos/gen/auth/trusted_device"
)

type TrustedDeviceHandler struct {
	trusteddevicepb.UnimplementedTrustedDeviceServiceServer
	trustedDeviceAppSvc *trustedDeviceApp.Service
}

func NewTrustedDeviceHandler(trustedDeviceAppSvc *trustedDeviceApp.Service) *TrustedDeviceHandler {
	return &TrustedDeviceHandler{
		trustedDeviceAppSvc: trustedDeviceAppSvc,
	}
}

// GetTrustedDeviceByID 根据ID获取受信任设备
func (h *TrustedDeviceHandler) GetTrustedDeviceByID(
	ctx context.Context,
	req *trusteddevicepb.GetTrustedDeviceByIDRequest,
) (*trusteddevicepb.GetTrustedDeviceByIDResponse, error) {
	return nil, errx.FailedPrecond("UNIMPLEMENTED", "method GetTrustedDeviceByID not implemented")
}
