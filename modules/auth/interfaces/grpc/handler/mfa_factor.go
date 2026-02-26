package handler

import (
	"context"

	mfaFactorApp "nfxidentity/modules/auth/application/mfa_factors"
	"nfxidentity/pkgs/errx"
	mfafactorpb "nfxidentity/protos/gen/auth/mfa_factor"
)

type MFAFactorHandler struct {
	mfafactorpb.UnimplementedMfaFactorServiceServer
	mfaFactorAppSvc *mfaFactorApp.Service
}

func NewMFAFactorHandler(mfaFactorAppSvc *mfaFactorApp.Service) *MFAFactorHandler {
	return &MFAFactorHandler{
		mfaFactorAppSvc: mfaFactorAppSvc,
	}
}

// GetMfaFactorByID 根据ID获取MFA因子
func (h *MFAFactorHandler) GetMfaFactorByID(ctx context.Context, req *mfafactorpb.GetMfaFactorByIDRequest) (*mfafactorpb.GetMfaFactorByIDResponse, error) {
	return nil, errx.FailedPrecond("UNIMPLEMENTED", "method GetMfaFactorByID not implemented")
}
