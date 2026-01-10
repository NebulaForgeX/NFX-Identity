package handler

import (
	"context"

	mfaFactorApp "nfxid/modules/auth/application/mfa_factors"
	mfafactorpb "nfxid/protos/gen/auth/mfa_factor"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	return nil, status.Errorf(codes.Unimplemented, "method GetMfaFactorByID not implemented")
}
