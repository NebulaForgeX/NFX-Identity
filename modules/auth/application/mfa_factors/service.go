package mfa_factors

import (
	mfaFactorDomain "nfxid/modules/auth/domain/mfa_factors"
)

type Service struct {
	mfaFactorRepo *mfaFactorDomain.Repo
}

func NewService(
	mfaFactorRepo *mfaFactorDomain.Repo,
) *Service {
	return &Service{
		mfaFactorRepo: mfaFactorRepo,
	}
}
