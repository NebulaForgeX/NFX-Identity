package hash_chain_checkpoints

import (
	"context"
	hashChainCheckpointCommands "nfxid/modules/audit/application/hash_chain_checkpoints/commands"
)

// DeleteHashChainCheckpoint 删除哈希链检查点
func (s *Service) DeleteHashChainCheckpoint(ctx context.Context, cmd hashChainCheckpointCommands.DeleteHashChainCheckpointCmd) error {
	// Delete from repository (hard delete)
	return s.hashChainCheckpointRepo.Delete.ByID(ctx, cmd.HashChainCheckpointID)
}
