package handler

import (
	"context"

	rateLimitApp "nfxid/modules/clients/application/rate_limits"
	ratelimitpb "nfxid/protos/gen/clients/rate_limit"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	return nil, status.Errorf(codes.Unimplemented, "method GetRateLimitByID not implemented")
}
