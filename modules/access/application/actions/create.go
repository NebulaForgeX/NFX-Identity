package actions

import (
	"context"

	actionCommands "nfxid/modules/access/application/actions/commands"
	actionDomain "nfxid/modules/access/domain/actions"

	"github.com/google/uuid"
)

func (s *Service) CreateAction(ctx context.Context, cmd actionCommands.CreateActionCmd) (uuid.UUID, error) {
	if exists, _ := s.actionRepo.Check.ByKey(ctx, cmd.Key); exists {
		return uuid.Nil, actionDomain.ErrActionKeyExists
	}
	action, err := actionDomain.NewAction(actionDomain.NewActionParams{
		Key:         cmd.Key,
		Service:     cmd.Service,
		Status:      cmd.Status,
		Name:        cmd.Name,
		Description: cmd.Description,
		IsSystem:    cmd.IsSystem,
	})
	if err != nil {
		return uuid.Nil, err
	}
	if err := s.actionRepo.Create.New(ctx, action); err != nil {
		return uuid.Nil, err
	}
	return action.ID(), nil
}
