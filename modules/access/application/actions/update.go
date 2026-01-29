package actions

import (
	"context"

	actionCommands "nfxid/modules/access/application/actions/commands"
)

func (s *Service) UpdateAction(ctx context.Context, cmd actionCommands.UpdateActionCmd) error {
	action, err := s.actionRepo.Get.ByID(ctx, cmd.ActionID)
	if err != nil {
		return err
	}
	if err := action.Update(cmd.Key, cmd.Service, cmd.Status, cmd.Name, cmd.Description); err != nil {
		return err
	}
	return s.actionRepo.Update.Generic(ctx, action)
}
