package respdto

import (
	"time"

	actorSnapshotAppResult "nfxid/modules/audit/application/actor_snapshots/results"

	"github.com/google/uuid"
)

type ActorSnapshotDTO struct {
	ID           uuid.UUID              `json:"id"`
	ActorType    string                 `json:"actor_type"`
	ActorID      uuid.UUID              `json:"actor_id"`
	DisplayName  *string                `json:"display_name,omitempty"`
	Email        *string                `json:"email,omitempty"`
	ClientName   *string                `json:"client_name,omitempty"`
	TenantID     *uuid.UUID             `json:"tenant_id,omitempty"`
	SnapshotAt   time.Time              `json:"snapshot_at"`
	SnapshotData map[string]interface{} `json:"snapshot_data,omitempty"`
	CreatedAt    time.Time              `json:"created_at"`
}

// ActorSnapshotROToDTO converts application ActorSnapshotRO to response DTO
func ActorSnapshotROToDTO(v *actorSnapshotAppResult.ActorSnapshotRO) *ActorSnapshotDTO {
	if v == nil {
		return nil
	}

	return &ActorSnapshotDTO{
		ID:           v.ID,
		ActorType:    string(v.ActorType),
		ActorID:      v.ActorID,
		DisplayName:  v.DisplayName,
		Email:        v.Email,
		ClientName:   v.ClientName,
		TenantID:     v.TenantID,
		SnapshotAt:   v.SnapshotAt,
		SnapshotData: v.SnapshotData,
		CreatedAt:    v.CreatedAt,
	}
}
