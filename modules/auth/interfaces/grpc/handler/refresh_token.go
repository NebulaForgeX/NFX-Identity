package handler

import (
	"context"

	refreshTokenApp "nfxidentity/modules/auth/application/refresh_tokens"
	"nfxidentity/pkgs/errx"
	refreshtokenpb "nfxidentity/protos/gen/auth/refresh_token"
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
func (h *RefreshTokenHandler) GetRefreshTokenByID(
	ctx context.Context,
	req *refreshtokenpb.GetRefreshTokenByIDRequest,
) (*refreshtokenpb.GetRefreshTokenByIDResponse, error) {
	return nil, errx.FailedPrecond("UNIMPLEMENTED", "method GetRefreshTokenByID not implemented")
}
