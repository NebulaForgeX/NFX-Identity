package reqdto

import (
	actorSnapshotAppCommands "nfxid/modules/audit/application/actor_snapshots/commands"
	actorSnapshotDomain "nfxid/modules/audit/domain/actor_snapshots"

	"github.com/google/uuid"
)

type ActorSnapshotCreateRequestDTO struct {
	ActorType    string                 `json:"actor_type" validate:"required"`
	ActorID      uuid.UUID              `json:"actor_id" validate:"required"`
	DisplayName  *string                `json:"display_name,omitempty"`
	Email        *string                `json:"email,omitempty"`
	ClientName   *string                `json:"client_name,omitempty"`
	TenantID     *uuid.UUID             `json:"tenant_id,omitempty"`
	SnapshotAt   string                 `json:"snapshot_at" validate:"required"`
	SnapshotData map[string]interface{} `json:"snapshot_data,omitempty"`
}

type ActorSnapshotByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

func (r *ActorSnapshotCreateRequestDTO) ToCreateCmd() actorSnapshotAppCommands.CreateActorSnapshotCmd {
	cmd := actorSnapshotAppCommands.CreateActorSnapshotCmd{
		ActorID:      r.ActorID,
		DisplayName:  r.DisplayName,
		Email:        r.Email,
		ClientName:   r.ClientName,
		TenantID:     r.TenantID,
		SnapshotAt:   r.SnapshotAt,
		SnapshotData: r.SnapshotData,
	}

	if r.ActorType != "" {
		cmd.ActorType = actorSnapshotDomain.ActorType(r.ActorType)
	}

	return cmd
}
