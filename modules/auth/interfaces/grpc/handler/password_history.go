package handler

import (
	"context"

	passwordHistoryApp "nfxid/modules/auth/application/password_history"
	passwordhistorypb "nfxid/protos/gen/auth/password_history"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
func (h *PasswordHistoryHandler) GetPasswordHistoryByID(ctx context.Context, req *passwordhistorypb.GetPasswordHistoryByIDRequest) (*passwordhistorypb.GetPasswordHistoryByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPasswordHistoryByID not implemented")
}
