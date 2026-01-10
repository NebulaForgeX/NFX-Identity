package user_emails

import (
	"context"
	userEmailCommands "nfxid/modules/directory/application/user_emails/commands"
	userEmailDomain "nfxid/modules/directory/domain/user_emails"

	"github.com/google/uuid"
)

// CreateUserEmail 创建用户邮箱
func (s *Service) CreateUserEmail(ctx context.Context, cmd userEmailCommands.CreateUserEmailCmd) (uuid.UUID, error) {
	// Check if email already exists
	if exists, _ := s.userEmailRepo.Check.ByEmail(ctx, cmd.Email); exists {
		return uuid.Nil, userEmailDomain.ErrEmailAlreadyExists
	}

	// Create domain entity
	userEmail, err := userEmailDomain.NewUserEmail(userEmailDomain.NewUserEmailParams{
		UserID:            cmd.UserID,
		Email:             cmd.Email,
		IsPrimary:         cmd.IsPrimary,
		IsVerified:        cmd.IsVerified,
		VerificationToken: cmd.VerificationToken,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.userEmailRepo.Create.New(ctx, userEmail); err != nil {
		return uuid.Nil, err
	}

	return userEmail.ID(), nil
}
