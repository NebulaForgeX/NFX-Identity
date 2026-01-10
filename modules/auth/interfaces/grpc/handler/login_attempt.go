package handler

import (
	"context"

	loginAttemptApp "nfxid/modules/auth/application/login_attempts"
	loginattemptpb "nfxid/protos/gen/auth/login_attempt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
func (h *LoginAttemptHandler) GetLoginAttemptByID(ctx context.Context, req *loginattemptpb.GetLoginAttemptByIDRequest) (*loginattemptpb.GetLoginAttemptByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoginAttemptByID not implemented")
}
