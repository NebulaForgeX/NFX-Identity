package password_history

import (
	"context"
	passwordHistoryCommands "nfxid/modules/auth/application/password_history/commands"
	passwordHistoryDomain "nfxid/modules/auth/domain/password_history"

	"github.com/google/uuid"
)

// CreatePasswordHistory 创建密码历史
func (s *Service) CreatePasswordHistory(ctx context.Context, cmd passwordHistoryCommands.CreatePasswordHistoryCmd) (uuid.UUID, error) {
	// Create domain entity
	passwordHistory, err := passwordHistoryDomain.NewPasswordHistory(passwordHistoryDomain.NewPasswordHistoryParams{
		UserID:       cmd.UserID,
		TenantID:     cmd.TenantID,
		PasswordHash: cmd.PasswordHash,
		HashAlg:      cmd.HashAlg,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.passwordHistoryRepo.Create.New(ctx, passwordHistory); err != nil {
		return uuid.Nil, err
	}

	return passwordHistory.ID(), nil
}
