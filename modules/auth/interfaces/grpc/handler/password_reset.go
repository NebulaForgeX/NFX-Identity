package handler

import (
	"context"

	passwordResetApp "nfxid/modules/auth/application/password_resets"
	passwordresetpb "nfxid/protos/gen/auth/password_reset"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PasswordResetHandler struct {
	passwordresetpb.UnimplementedPasswordResetServiceServer
	passwordResetAppSvc *passwordResetApp.Service
}

func NewPasswordResetHandler(passwordResetAppSvc *passwordResetApp.Service) *PasswordResetHandler {
	return &PasswordResetHandler{
		passwordResetAppSvc: passwordResetAppSvc,
	}
}

// GetPasswordResetByID 根据ID获取密码重置
func (h *PasswordResetHandler) GetPasswordResetByID(ctx context.Context, req *passwordresetpb.GetPasswordResetByIDRequest) (*passwordresetpb.GetPasswordResetByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPasswordResetByID not implemented")
}
