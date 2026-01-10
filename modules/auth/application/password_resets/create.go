package password_resets

import (
	"context"
	"time"
	passwordResetCommands "nfxid/modules/auth/application/password_resets/commands"
	passwordResetDomain "nfxid/modules/auth/domain/password_resets"

	"github.com/google/uuid"
)

// CreatePasswordReset 创建密码重置
func (s *Service) CreatePasswordReset(ctx context.Context, cmd passwordResetCommands.CreatePasswordResetCmd) (uuid.UUID, error) {
	// Parse expires at
	expiresAt, err := time.Parse(time.RFC3339, cmd.ExpiresAt)
	if err != nil {
		return uuid.Nil, err
	}

	// Create domain entity
	passwordReset, err := passwordResetDomain.NewPasswordReset(passwordResetDomain.NewPasswordResetParams{
		ResetID:     cmd.ResetID,
		TenantID:    cmd.TenantID,
		UserID:      cmd.UserID,
		Delivery:    cmd.Delivery,
		CodeHash:    cmd.CodeHash,
		ExpiresAt:   expiresAt,
		RequestedIP: cmd.RequestedIP,
		UAHash:      cmd.UAHash,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.passwordResetRepo.Create.New(ctx, passwordReset); err != nil {
		return uuid.Nil, err
	}

	return passwordReset.ID(), nil
}
