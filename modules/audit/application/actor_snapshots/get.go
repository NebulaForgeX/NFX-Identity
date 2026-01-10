package actor_snapshots

import (
	"context"
	actorSnapshotResult "nfxid/modules/audit/application/actor_snapshots/results"

	"github.com/google/uuid"
)

// GetActorSnapshot 根据ID获取参与者快照
func (s *Service) GetActorSnapshot(ctx context.Context, actorSnapshotID uuid.UUID) (actorSnapshotResult.ActorSnapshotRO, error) {
	domainEntity, err := s.actorSnapshotRepo.Get.ByID(ctx, actorSnapshotID)
	if err != nil {
		return actorSnapshotResult.ActorSnapshotRO{}, err
	}
	return actorSnapshotResult.ActorSnapshotMapper(domainEntity), nil
}
