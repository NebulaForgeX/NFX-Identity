package mfa_factors

import (
	"context"
	mfaFactorResult "nfxid/modules/auth/application/mfa_factors/results"

	"github.com/google/uuid"
)

// GetMFAFactor 根据ID获取MFA因子
func (s *Service) GetMFAFactor(ctx context.Context, mfaFactorID uuid.UUID) (mfaFactorResult.MFAFactorRO, error) {
	domainEntity, err := s.mfaFactorRepo.Get.ByID(ctx, mfaFactorID)
	if err != nil {
		return mfaFactorResult.MFAFactorRO{}, err
	}
	return mfaFactorResult.MFAFactorMapper(domainEntity), nil
}

// GetMFAFactorByFactorID 根据FactorID获取MFA因子
func (s *Service) GetMFAFactorByFactorID(ctx context.Context, factorID string) (mfaFactorResult.MFAFactorRO, error) {
	domainEntity, err := s.mfaFactorRepo.Get.ByFactorID(ctx, factorID)
	if err != nil {
		return mfaFactorResult.MFAFactorRO{}, err
	}
	return mfaFactorResult.MFAFactorMapper(domainEntity), nil
}
