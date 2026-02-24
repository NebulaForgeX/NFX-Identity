package handler

import (
	"context"

	passwordHistoryApp "nfxid/modules/auth/application/password_history"
	"nfxid/pkgs/errx"
	passwordhistorypb "nfxid/protos/gen/auth/password_history"
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
