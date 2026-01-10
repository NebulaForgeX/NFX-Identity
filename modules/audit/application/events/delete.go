package events

import (
	"context"
	eventCommands "nfxid/modules/audit/application/events/commands"
)

// DeleteEvent 删除事件
func (s *Service) DeleteEvent(ctx context.Context, cmd eventCommands.DeleteEventCmd) error {
	// Delete from repository (hard delete)
	return s.eventRepo.Delete.ByID(ctx, cmd.EventID)
}
