package actions

import (
	"context"

	actionCommands "nfxid/modules/access/application/actions/commands"
)

func (s *Service) DeleteAction(ctx context.Context, cmd actionCommands.DeleteActionCmd) error {
	action, err := s.actionRepo.Get.ByID(ctx, cmd.ActionID)
	if err != nil {
		return err
	}
	if err := action.Delete(); err != nil {
		return err
	}
	return s.actionRepo.Update.Generic(ctx, action)
}
