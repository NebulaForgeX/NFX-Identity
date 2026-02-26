package handler

import (
	"context"

	passwordHistoryApp "nfxidentity/modules/auth/application/password_history"
	"nfxidentity/pkgs/errx"
	passwordhistorypb "nfxidentity/protos/gen/auth/password_history"
)

type PasswordHistoryHandler struct {
	passwordhistorypb.UnimplementedPasswordHistoryServiceServer
	passwordHistoryAppSvc *passwordHistoryApp.Service
}

func NewPasswordHistoryHandler(passwordHistoryAppSvc *passwordHistoryApp.Service) *PasswordHistoryHandler {
	return &PasswordHistoryHandler{
		passwordHistoryAppSvc: passwordHistoryAppSvc,
	}
}

// GetPasswordHistoryByID 根据ID获取密码历史
func (h *PasswordHistoryHandler) GetPasswordHistoryByID(
	ctx context.Context,
	req *passwordhistorypb.GetPasswordHistoryByIDRequest,
) (*passwordhistorypb.GetPasswordHistoryByIDResponse, error) {
	return nil, errx.FailedPrecond("UNIMPLEMENTED", "method GetPasswordHistoryByID not implemented")
}
