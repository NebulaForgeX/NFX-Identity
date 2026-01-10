package commands

import (
	"nfxid/modules/audit/domain/actor_snapshots"

	"github.com/google/uuid"
)

// CreateActorSnapshotCmd 创建参与者快照命令
type CreateActorSnapshotCmd struct {
	ActorType    actor_snapshots.ActorType
	ActorID      uuid.UUID
	DisplayName  *string
	Email        *string
	ClientName   *string
	TenantID     *uuid.UUID
	SnapshotAt   string
	SnapshotData map[string]interface{}
}

// DeleteActorSnapshotCmd 删除参与者快照命令
type DeleteActorSnapshotCmd struct {
	ActorSnapshotID uuid.UUID
}
