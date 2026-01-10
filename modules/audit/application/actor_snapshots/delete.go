package actor_snapshots

import (
	"context"
	actorSnapshotCommands "nfxid/modules/audit/application/actor_snapshots/commands"
)

// DeleteActorSnapshot 删除参与者快照
func (s *Service) DeleteActorSnapshot(ctx context.Context, cmd actorSnapshotCommands.DeleteActorSnapshotCmd) error {
	// Delete from repository (hard delete)
	return s.actorSnapshotRepo.Delete.ByID(ctx, cmd.ActorSnapshotID)
}
