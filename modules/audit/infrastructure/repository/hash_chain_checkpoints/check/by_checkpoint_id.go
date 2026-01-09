package check

import (
	"context"
	"nfxid/modules/audit/infrastructure/rdb/models"
)

// ByCheckpointID 根据 CheckpointID 检查 HashChainCheckpoint 是否存在，实现 hash_chain_checkpoints.Check 接口
func (h *Handler) ByCheckpointID(ctx context.Context, checkpointID string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.HashChainCheckpoint{}).
		Where("checkpoint_id = ?", checkpointID).
		Count(&count).Error
	return count > 0, err
}
