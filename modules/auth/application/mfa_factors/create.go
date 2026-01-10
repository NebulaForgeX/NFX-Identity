package mfa_factors

import (
	"context"
	mfaFactorCommands "nfxid/modules/auth/application/mfa_factors/commands"
	mfaFactorDomain "nfxid/modules/auth/domain/mfa_factors"

	"github.com/google/uuid"
)

// CreateMFAFactor 创建MFA因子
func (s *Service) CreateMFAFactor(ctx context.Context, cmd mfaFactorCommands.CreateMFAFactorCmd) (uuid.UUID, error) {
	// Create domain entity
	mfaFactor, err := mfaFactorDomain.NewMFAFactor(mfaFactorDomain.NewMFAFactorParams{
		FactorID:          cmd.FactorID,
		TenantID:          cmd.TenantID,
		UserID:            cmd.UserID,
		Type:              cmd.Type,
		SecretEncrypted:   cmd.SecretEncrypted,
		Phone:             cmd.Phone,
		Email:             cmd.Email,
		Name:              cmd.Name,
		Enabled:           cmd.Enabled,
		RecoveryCodesHash: cmd.RecoveryCodesHash,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.mfaFactorRepo.Create.New(ctx, mfaFactor); err != nil {
		return uuid.Nil, err
	}

	return mfaFactor.ID(), nil
}
