package actor_snapshots

import (
	actorSnapshotDomain "nfxid/modules/audit/domain/actor_snapshots"
)

type Service struct {
	actorSnapshotRepo *actorSnapshotDomain.Repo
}

func NewService(
	actorSnapshotRepo *actorSnapshotDomain.Repo,
) *Service {
	return &Service{
		actorSnapshotRepo: actorSnapshotRepo,
	}
}
