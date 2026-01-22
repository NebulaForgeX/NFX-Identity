package user_credentials

import (
	"context"
	userCredentialCommands "nfxid/modules/auth/application/user_credentials/commands"
	userCredentialDomain "nfxid/modules/auth/domain/user_credentials"

	"github.com/google/uuid"
)

// CreateUserCredential 创建用户凭证
func (s *Service) CreateUserCredential(ctx context.Context, cmd userCredentialCommands.CreateUserCredentialCmd) (uuid.UUID, error) {
	// Create domain entity
	userCredential, err := userCredentialDomain.NewUserCredential(userCredentialDomain.NewUserCredentialParams{
		UserID:             cmd.UserID,
		CredentialType:     cmd.CredentialType,
		PasswordHash:       cmd.PasswordHash,
		HashAlg:            cmd.HashAlg,
		HashParams:         cmd.HashParams,
		Status:             cmd.Status,
		MustChangePassword: cmd.MustChangePassword,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.userCredentialRepo.Create.New(ctx, userCredential); err != nil {
		return uuid.Nil, err
	}

	return userCredential.ID(), nil
}
