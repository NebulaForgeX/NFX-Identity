package hash_chain_checkpoints

import (
	"time"

	"github.com/google/uuid"
)

type NewHashChainCheckpointParams struct {
	CheckpointID       string
	TenantID           *uuid.UUID
	PartitionDate      time.Time
	CheckpointHash     string
	PrevCheckpointHash *string
	EventCount         int
	FirstEventID       *string
	LastEventID        *string
	CreatedBy          *string
}

func NewHashChainCheckpoint(p NewHashChainCheckpointParams) (*HashChainCheckpoint, error) {
	if err := validateHashChainCheckpointParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewHashChainCheckpointFromState(HashChainCheckpointState{
		ID:                 id,
		CheckpointID:       p.CheckpointID,
		TenantID:           p.TenantID,
		PartitionDate:      p.PartitionDate,
		CheckpointHash:     p.CheckpointHash,
		PrevCheckpointHash: p.PrevCheckpointHash,
		EventCount:         p.EventCount,
		FirstEventID:       p.FirstEventID,
		LastEventID:        p.LastEventID,
		CreatedAt:          now,
		CreatedBy:          p.CreatedBy,
	}), nil
}

func NewHashChainCheckpointFromState(st HashChainCheckpointState) *HashChainCheckpoint {
	return &HashChainCheckpoint{state: st}
}

func validateHashChainCheckpointParams(p NewHashChainCheckpointParams) error {
	if p.CheckpointID == "" {
		return ErrCheckpointIDRequired
	}
	if p.PartitionDate.IsZero() {
		return ErrPartitionDateRequired
	}
	if p.CheckpointHash == "" {
		return ErrCheckpointHashRequired
	}
	return nil
}
