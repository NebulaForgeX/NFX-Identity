package hash_chain_checkpoints

import (
	"context"
	hashChainCheckpointResult "nfxid/modules/audit/application/hash_chain_checkpoints/results"

	"github.com/google/uuid"
)

// GetHashChainCheckpoint 根据ID获取哈希链检查点
func (s *Service) GetHashChainCheckpoint(ctx context.Context, hashChainCheckpointID uuid.UUID) (hashChainCheckpointResult.HashChainCheckpointRO, error) {
	domainEntity, err := s.hashChainCheckpointRepo.Get.ByID(ctx, hashChainCheckpointID)
	if err != nil {
		return hashChainCheckpointResult.HashChainCheckpointRO{}, err
	}
	return hashChainCheckpointResult.HashChainCheckpointMapper(domainEntity), nil
}
