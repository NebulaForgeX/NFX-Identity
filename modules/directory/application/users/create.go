package users

import (
	"context"
	userCommands "nfxid/modules/directory/application/users/commands"
	userDomain "nfxid/modules/directory/domain/users"

	"github.com/google/uuid"
)

// CreateUser 创建用户
func (s *Service) CreateUser(ctx context.Context, cmd userCommands.CreateUserCmd) (uuid.UUID, error) {
	// Check if username already exists
	if exists, _ := s.userRepo.Check.ByUsername(ctx, cmd.Username); exists {
		return uuid.Nil, userDomain.ErrUsernameAlreadyExists
	}

	// Create domain entity
	user, err := userDomain.NewUser(userDomain.NewUserParams{
		Username:   cmd.Username,
		Status:     cmd.Status,
		IsVerified: cmd.IsVerified,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.userRepo.Create.New(ctx, user); err != nil {
		return uuid.Nil, err
	}

	return user.ID(), nil
}
