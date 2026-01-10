package results

import (
	"time"

	"nfxid/modules/audit/domain/hash_chain_checkpoints"

	"github.com/google/uuid"
)

type HashChainCheckpointRO struct {
	ID                 uuid.UUID
	CheckpointID       string
	TenantID           *uuid.UUID
	PartitionDate      time.Time
	CheckpointHash     string
	PrevCheckpointHash *string
	EventCount         int
	FirstEventID       *string
	LastEventID        *string
	CreatedAt          time.Time
	CreatedBy          *string
}

// HashChainCheckpointMapper 将 Domain HashChainCheckpoint 转换为 Application HashChainCheckpointRO
func HashChainCheckpointMapper(hcc *hash_chain_checkpoints.HashChainCheckpoint) HashChainCheckpointRO {
	if hcc == nil {
		return HashChainCheckpointRO{}
	}

	return HashChainCheckpointRO{
		ID:                 hcc.ID(),
		CheckpointID:       hcc.CheckpointID(),
		TenantID:           hcc.TenantID(),
		PartitionDate:      hcc.PartitionDate(),
		CheckpointHash:     hcc.CheckpointHash(),
		PrevCheckpointHash: hcc.PrevCheckpointHash(),
		EventCount:         hcc.EventCount(),
		FirstEventID:       hcc.FirstEventID(),
		LastEventID:        hcc.LastEventID(),
		CreatedAt:          hcc.CreatedAt(),
		CreatedBy:          hcc.CreatedBy(),
	}
}
