package handler

import (
	"context"

	ipAllowlistApp "nfxid/modules/clients/application/ip_allowlist"
	ipallowlistpb "nfxid/protos/gen/clients/ip_allowlist"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	return nil, status.Errorf(codes.Unimplemented, "method GetIpAllowlistByID not implemented")
}
