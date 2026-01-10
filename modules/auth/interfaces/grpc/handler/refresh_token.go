package handler

import (
	"context"

	refreshTokenApp "nfxid/modules/auth/application/refresh_tokens"
	refreshtokenpb "nfxid/protos/gen/auth/refresh_token"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RefreshTokenHandler struct {
	refreshtokenpb.UnimplementedRefreshTokenServiceServer
	refreshTokenAppSvc *refreshTokenApp.Service
}

func NewRefreshTokenHandler(refreshTokenAppSvc *refreshTokenApp.Service) *RefreshTokenHandler {
	return &RefreshTokenHandler{
		refreshTokenAppSvc: refreshTokenAppSvc,
	}
}

// GetRefreshTokenByID 根据ID获取刷新令牌
func (h *RefreshTokenHandler) GetRefreshTokenByID(ctx context.Context, req *refreshtokenpb.GetRefreshTokenByIDRequest) (*refreshtokenpb.GetRefreshTokenByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRefreshTokenByID not implemented")
}
