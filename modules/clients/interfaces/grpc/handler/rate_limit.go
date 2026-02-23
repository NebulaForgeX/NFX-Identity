package handler

import (
	"context"

	rateLimitApp "nfxid/modules/clients/application/rate_limits"
	ratelimitpb "nfxid/protos/gen/clients/rate_limit"
	"nfxid/pkgs/errx"
)

type RateLimitHandler struct {
	ratelimitpb.UnimplementedRateLimitServiceServer
	rateLimitAppSvc *rateLimitApp.Service
}

func NewRateLimitHandler(rateLimitAppSvc *rateLimitApp.Service) *RateLimitHandler {
	return &RateLimitHandler{
		rateLimitAppSvc: rateLimitAppSvc,
	}
}

// GetRateLimitByID 根据ID获取Rate Limit
func (h *RateLimitHandler) GetRateLimitByID(ctx context.Context, req *ratelimitpb.GetRateLimitByIDRequest) (*ratelimitpb.GetRateLimitByIDResponse, error) {
	return nil, errx.FailedPrecond("UNIMPLEMENTED", "method GetRateLimitByID not implemented")
}
