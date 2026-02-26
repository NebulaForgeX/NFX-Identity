package handler

import (
	"context"

	passwordResetApp "nfxidentity/modules/auth/application/password_resets"
	"nfxidentity/pkgs/errx"
	passwordresetpb "nfxidentity/protos/gen/auth/password_reset"
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
func (h *PasswordResetHandler) GetPasswordResetByID(
	ctx context.Context,
	req *passwordresetpb.GetPasswordResetByIDRequest,
) (*passwordresetpb.GetPasswordResetByIDResponse, error) {
	return nil, errx.FailedPrecond("UNIMPLEMENTED", "method GetPasswordResetByID not implemented")
}
