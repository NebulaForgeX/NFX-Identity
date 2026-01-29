package action_requirements

import (
	"context"

	arCommands "nfxid/modules/access/application/action_requirements/commands"
)

func (s *Service) DeleteActionRequirement(ctx context.Context, cmd arCommands.DeleteActionRequirementCmd) error {
	return s.arRepo.Delete.ByID(ctx, cmd.ActionRequirementID)
}

func (s *Service) DeleteByActionIDAndPermissionID(ctx context.Context, cmd arCommands.DeleteByActionIDAndPermissionIDCmd) error {
	return s.arRepo.Delete.ByActionIDAndPermissionID(ctx, cmd.ActionID, cmd.PermissionID)
}

func (s *Service) DeleteByActionID(ctx context.Context, cmd arCommands.DeleteByActionIDCmd) error {
	return s.arRepo.Delete.ByActionID(ctx, cmd.ActionID)
}
