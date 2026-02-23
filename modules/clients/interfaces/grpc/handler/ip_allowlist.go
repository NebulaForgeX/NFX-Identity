package handler

import (
	"context"

	ipAllowlistApp "nfxid/modules/clients/application/ip_allowlist"
	ipallowlistpb "nfxid/protos/gen/clients/ip_allowlist"
	"nfxid/pkgs/errx"
)

type IPAllowlistHandler struct {
	ipallowlistpb.UnimplementedIpAllowlistServiceServer
	ipAllowlistAppSvc *ipAllowlistApp.Service
}

func NewIPAllowlistHandler(ipAllowlistAppSvc *ipAllowlistApp.Service) *IPAllowlistHandler {
	return &IPAllowlistHandler{
		ipAllowlistAppSvc: ipAllowlistAppSvc,
	}
}

// GetIpAllowlistByID 根据ID获取IP Allowlist
func (h *IPAllowlistHandler) GetIpAllowlistByID(ctx context.Context, req *ipallowlistpb.GetIpAllowlistByIDRequest) (*ipallowlistpb.GetIpAllowlistByIDResponse, error) {
	return nil, errx.FailedPrecond("UNIMPLEMENTED", "method GetIpAllowlistByID not implemented")
}
