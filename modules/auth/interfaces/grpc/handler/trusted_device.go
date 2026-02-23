package handler

import (
	"context"

	trustedDeviceApp "nfxid/modules/auth/application/trusted_devices"
	trusteddevicepb "nfxid/protos/gen/auth/trusted_device"
	"nfxid/pkgs/errx"
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
func (h *TrustedDeviceHandler) GetTrustedDeviceByID(ctx context.Context, req *trusteddevicepb.GetTrustedDeviceByIDRequest) (*trusteddevicepb.GetTrustedDeviceByIDResponse, error) {
	return nil, errx.FailedPrecond("UNIMPLEMENTED", "method GetTrustedDeviceByID not implemented")
}
