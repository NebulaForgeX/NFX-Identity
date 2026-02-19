package reqdto

import (
	hashChainCheckpointAppCommands "nfxid/modules/audit/application/hash_chain_checkpoints/commands"

	"github.com/google/uuid"
)

type HashChainCheckpointCreateRequestDTO struct {
	CheckpointID       string     `json:"checkpoint_id" validate:"required"`
	TenantID           *uuid.UUID `json:"tenant_id,omitempty"`
	PartitionDate      string     `json:"partition_date" validate:"required"`
	CheckpointHash     string     `json:"checkpoint_hash" validate:"required"`
	PrevCheckpointHash *string    `json:"prev_checkpoint_hash,omitempty"`
	EventCount         int        `json:"event_count" validate:"required"`
	FirstEventID       *string    `json:"first_event_id,omitempty"`
	LastEventID        *string    `json:"last_event_id,omitempty"`
	CreatedBy          *string    `json:"created_by,omitempty"`
}

type HashChainCheckpointByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

func (r *HashChainCheckpointCreateRequestDTO) ToCreateCmd() hashChainCheckpointAppCommands.CreateHashChainCheckpointCmd {
	return hashChainCheckpointAppCommands.CreateHashChainCheckpointCmd{
		CheckpointID:       r.CheckpointID,
		TenantID:           r.TenantID,
		PartitionDate:      r.PartitionDate,
		CheckpointHash:     r.CheckpointHash,
		PrevCheckpointHash: r.PrevCheckpointHash,
		EventCount:         r.EventCount,
		FirstEventID:       r.FirstEventID,
		LastEventID:        r.LastEventID,
		CreatedBy:          r.CreatedBy,
	}
}
