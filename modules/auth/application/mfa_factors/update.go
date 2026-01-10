package mfa_factors

import (
	"context"
	mfaFactorCommands "nfxid/modules/auth/application/mfa_factors/commands"
)

// UpdateMFAFactor 更新MFA因子
func (s *Service) UpdateMFAFactor(ctx context.Context, cmd mfaFactorCommands.UpdateMFAFactorCmd) error {
	// Get domain entity
	mfaFactor, err := s.mfaFactorRepo.Get.ByFactorID(ctx, cmd.FactorID)
	if err != nil {
		return err
	}

	// Update name if provided
	if cmd.Name != nil {
		if err := mfaFactor.UpdateName(cmd.Name); err != nil {
			return err
		}
	}

	// Update enabled status if needed
	if cmd.Enabled {
		if err := mfaFactor.Enable(); err != nil {
			return err
		}
	} else {
		if err := mfaFactor.Disable(); err != nil {
			return err
		}
	}

	// Save to repository
	return s.mfaFactorRepo.Update.Generic(ctx, mfaFactor)
}

// EnableMFAFactor 启用MFA因子
func (s *Service) EnableMFAFactor(ctx context.Context, cmd mfaFactorCommands.EnableMFAFactorCmd) error {
	// Get domain entity
	mfaFactor, err := s.mfaFactorRepo.Get.ByFactorID(ctx, cmd.FactorID)
	if err != nil {
		return err
	}

	// Enable domain entity
	if err := mfaFactor.Enable(); err != nil {
		return err
	}

	// Save to repository
	return s.mfaFactorRepo.Update.Enable(ctx, cmd.FactorID)
}

// DisableMFAFactor 禁用MFA因子
func (s *Service) DisableMFAFactor(ctx context.Context, cmd mfaFactorCommands.DisableMFAFactorCmd) error {
	// Get domain entity
	mfaFactor, err := s.mfaFactorRepo.Get.ByFactorID(ctx, cmd.FactorID)
	if err != nil {
		return err
	}

	// Disable domain entity
	if err := mfaFactor.Disable(); err != nil {
		return err
	}

	// Save to repository
	return s.mfaFactorRepo.Update.Disable(ctx, cmd.FactorID)
}

// UpdateLastUsed 更新最后使用时间
func (s *Service) UpdateLastUsed(ctx context.Context, cmd mfaFactorCommands.UpdateLastUsedCmd) error {
	// Get domain entity
	mfaFactor, err := s.mfaFactorRepo.Get.ByFactorID(ctx, cmd.FactorID)
	if err != nil {
		return err
	}

	// Update last used domain entity
	if err := mfaFactor.UpdateLastUsed(); err != nil {
		return err
	}

	// Save to repository
	return s.mfaFactorRepo.Update.UpdateLastUsed(ctx, cmd.FactorID)
}
