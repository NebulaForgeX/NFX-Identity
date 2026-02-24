package handler

import (
	"context"

	accountLockoutApp "nfxid/modules/auth/application/account_lockouts"
	"nfxid/pkgs/errx"
	accountlockoutpb "nfxid/protos/gen/auth/account_lockout"
)

type AccountLockoutHandler struct {
	accountlockoutpb.UnimplementedAccountLockoutServiceServer
	accountLockoutAppSvc *accountLockoutApp.Service
}

func NewAccountLockoutHandler(accountLockoutAppSvc *accountLockoutApp.Service) *AccountLockoutHandler {
	return &AccountLockoutHandler{
		accountLockoutAppSvc: accountLockoutAppSvc,
	}
}

// GetAccountLockoutByUserID 根据用户ID获取账户锁定
func (h *AccountLockoutHandler) GetAccountLockoutByUserID(
	ctx context.Context,
	req *accountlockoutpb.GetAccountLockoutByUserIDRequest,
) (*accountlockoutpb.GetAccountLockoutByUserIDResponse, error) {
	return nil, errx.FailedPrecond("UNIMPLEMENTED", "method GetAccountLockoutByUserID not implemented")
}

// BatchGetAccountLockouts 批量获取账户锁定
func (h *AccountLockoutHandler) BatchGetAccountLockouts(
	ctx context.Context,
	req *accountlockoutpb.BatchGetAccountLockoutsRequest,
) (*accountlockoutpb.BatchGetAccountLockoutsResponse, error) {
	return nil, errx.FailedPrecond("UNIMPLEMENTED", "method BatchGetAccountLockouts not implemented")
}
