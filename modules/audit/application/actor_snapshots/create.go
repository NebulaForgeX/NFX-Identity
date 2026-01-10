package actor_snapshots

import (
	"context"
	"time"
	actorSnapshotCommands "nfxid/modules/audit/application/actor_snapshots/commands"
	actorSnapshotDomain "nfxid/modules/audit/domain/actor_snapshots"

	"github.com/google/uuid"
)

// CreateActorSnapshot 创建参与者快照
func (s *Service) CreateActorSnapshot(ctx context.Context, cmd actorSnapshotCommands.CreateActorSnapshotCmd) (uuid.UUID, error) {
	// Parse snapshot at
	var snapshotAt time.Time
	if cmd.SnapshotAt != "" {
		parsed, err := time.Parse(time.RFC3339, cmd.SnapshotAt)
		if err != nil {
			return uuid.Nil, err
		}
		snapshotAt = parsed
	}

	// Create domain entity
	actorSnapshot, err := actorSnapshotDomain.NewActorSnapshot(actorSnapshotDomain.NewActorSnapshotParams{
		ActorType:    cmd.ActorType,
		ActorID:      cmd.ActorID,
		DisplayName:  cmd.DisplayName,
		Email:        cmd.Email,
		ClientName:   cmd.ClientName,
		TenantID:     cmd.TenantID,
		SnapshotAt:   snapshotAt,
		SnapshotData: cmd.SnapshotData,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.actorSnapshotRepo.Create.New(ctx, actorSnapshot); err != nil {
		return uuid.Nil, err
	}

	return actorSnapshot.ID(), nil
}
