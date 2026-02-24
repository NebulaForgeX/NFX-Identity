package handler

import (
	"context"

	loginAttemptApp "nfxid/modules/auth/application/login_attempts"
	"nfxid/pkgs/errx"
	loginattemptpb "nfxid/protos/gen/auth/login_attempt"
)

type LoginAttemptHandler struct {
	loginattemptpb.UnimplementedLoginAttemptServiceServer
	loginAttemptAppSvc *loginAttemptApp.Service
}

func NewLoginAttemptHandler(loginAttemptAppSvc *loginAttemptApp.Service) *LoginAttemptHandler {
	return &LoginAttemptHandler{
		loginAttemptAppSvc: loginAttemptAppSvc,
	}
}

// GetLoginAttemptByID 根据ID获取登录尝试
func (h *LoginAttemptHandler) GetLoginAttemptByID(
	ctx context.Context,
	req *loginattemptpb.GetLoginAttemptByIDRequest,
) (*loginattemptpb.GetLoginAttemptByIDResponse, error) {
	return nil, errx.FailedPrecond("UNIMPLEMENTED", "method GetLoginAttemptByID not implemented")
}
