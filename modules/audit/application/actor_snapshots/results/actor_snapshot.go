package results

import (
	"time"

	"nfxid/modules/audit/domain/actor_snapshots"

	"github.com/google/uuid"
)

type ActorSnapshotRO struct {
	ID           uuid.UUID
	ActorType    actor_snapshots.ActorType
	ActorID      uuid.UUID
	DisplayName  *string
	Email        *string
	ClientName   *string
	TenantID     *uuid.UUID
	SnapshotAt   time.Time
	SnapshotData map[string]interface{}
	CreatedAt    time.Time
}

// ActorSnapshotMapper 将 Domain ActorSnapshot 转换为 Application ActorSnapshotRO
func ActorSnapshotMapper(as *actor_snapshots.ActorSnapshot) ActorSnapshotRO {
	if as == nil {
		return ActorSnapshotRO{}
	}

	return ActorSnapshotRO{
		ID:           as.ID(),
		ActorType:    as.ActorType(),
		ActorID:      as.ActorID(),
		DisplayName:  as.DisplayName(),
		Email:        as.Email(),
		ClientName:   as.ClientName(),
		TenantID:     as.TenantID(),
		SnapshotAt:   as.SnapshotAt(),
		SnapshotData: as.SnapshotData(),
		CreatedAt:    as.CreatedAt(),
	}
}
