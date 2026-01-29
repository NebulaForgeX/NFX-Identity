package action_requirements

import (
	"context"

	arCommands "nfxid/modules/access/application/action_requirements/commands"
	arDomain "nfxid/modules/access/domain/action_requirements"

	"github.com/google/uuid"
)

func (s *Service) CreateActionRequirement(ctx context.Context, cmd arCommands.CreateActionRequirementCmd) (uuid.UUID, error) {
	exists, _ := s.arRepo.Check.ByActionIDAndPermissionID(ctx, cmd.ActionID, cmd.PermissionID)
	if exists {
		return uuid.Nil, arDomain.ErrActionRequirementAlreadyExists
	}
	ar, err := arDomain.NewActionRequirement(arDomain.NewActionRequirementParams{
		ActionID:     cmd.ActionID,
		PermissionID: cmd.PermissionID,
		GroupID:      cmd.GroupID,
	})
	if err != nil {
		return uuid.Nil, err
	}
	if err := s.arRepo.Create.New(ctx, ar); err != nil {
		return uuid.Nil, err
	}
	return ar.ID(), nil
}
