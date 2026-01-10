package mfa_factors

import (
	"context"
	mfaFactorCommands "nfxid/modules/auth/application/mfa_factors/commands"
)

// DeleteMFAFactor 删除MFA因子（软删除）
func (s *Service) DeleteMFAFactor(ctx context.Context, cmd mfaFactorCommands.DeleteMFAFactorCmd) error {
	// Get domain entity
	mfaFactor, err := s.mfaFactorRepo.Get.ByFactorID(ctx, cmd.FactorID)
	if err != nil {
		return err
	}

	// Delete domain entity (soft delete)
	if err := mfaFactor.Delete(); err != nil {
		return err
	}

	// Save to repository
	return s.mfaFactorRepo.Update.Generic(ctx, mfaFactor)
}
