package respdto

import (
	"time"

	hashChainCheckpointAppResult "nfxid/modules/audit/application/hash_chain_checkpoints/results"

	"github.com/google/uuid"
)

type HashChainCheckpointDTO struct {
	ID                 uuid.UUID  `json:"id"`
	CheckpointID       string      `json:"checkpoint_id"`
	TenantID           *uuid.UUID  `json:"tenant_id,omitempty"`
	PartitionDate      time.Time   `json:"partition_date"`
	CheckpointHash     string      `json:"checkpoint_hash"`
	PrevCheckpointHash *string     `json:"prev_checkpoint_hash,omitempty"`
	EventCount         int         `json:"event_count"`
	FirstEventID       *string     `json:"first_event_id,omitempty"`
	LastEventID        *string     `json:"last_event_id,omitempty"`
	CreatedAt          time.Time   `json:"created_at"`
	CreatedBy          *string     `json:"created_by,omitempty"`
}

// HashChainCheckpointROToDTO converts application HashChainCheckpointRO to response DTO
func HashChainCheckpointROToDTO(v *hashChainCheckpointAppResult.HashChainCheckpointRO) *HashChainCheckpointDTO {
	if v == nil {
		return nil
	}

	return &HashChainCheckpointDTO{
		ID:                 v.ID,
		CheckpointID:       v.CheckpointID,
		TenantID:           v.TenantID,
		PartitionDate:      v.PartitionDate,
		CheckpointHash:     v.CheckpointHash,
		PrevCheckpointHash: v.PrevCheckpointHash,
		EventCount:         v.EventCount,
		FirstEventID:       v.FirstEventID,
		LastEventID:        v.LastEventID,
		CreatedAt:          v.CreatedAt,
		CreatedBy:          v.CreatedBy,
	}
}
