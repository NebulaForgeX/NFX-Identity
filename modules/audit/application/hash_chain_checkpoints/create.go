package hash_chain_checkpoints

import (
	"context"
	"time"
	hashChainCheckpointCommands "nfxid/modules/audit/application/hash_chain_checkpoints/commands"
	hashChainCheckpointDomain "nfxid/modules/audit/domain/hash_chain_checkpoints"

	"github.com/google/uuid"
)

// CreateHashChainCheckpoint 创建哈希链检查点
func (s *Service) CreateHashChainCheckpoint(ctx context.Context, cmd hashChainCheckpointCommands.CreateHashChainCheckpointCmd) (uuid.UUID, error) {
	// Parse partition date
	partitionDate, err := time.Parse("2006-01-02", cmd.PartitionDate)
	if err != nil {
		return uuid.Nil, err
	}

	// Create domain entity
	hashChainCheckpoint, err := hashChainCheckpointDomain.NewHashChainCheckpoint(hashChainCheckpointDomain.NewHashChainCheckpointParams{
		CheckpointID:       cmd.CheckpointID,
		TenantID:           cmd.TenantID,
		PartitionDate:      partitionDate,
		CheckpointHash:     cmd.CheckpointHash,
		PrevCheckpointHash: cmd.PrevCheckpointHash,
		EventCount:         cmd.EventCount,
		FirstEventID:       cmd.FirstEventID,
		LastEventID:        cmd.LastEventID,
		CreatedBy:          cmd.CreatedBy,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.hashChainCheckpointRepo.Create.New(ctx, hashChainCheckpoint); err != nil {
		return uuid.Nil, err
	}

	return hashChainCheckpoint.ID(), nil
}
