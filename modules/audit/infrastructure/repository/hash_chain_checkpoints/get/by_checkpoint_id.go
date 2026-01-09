package get

import (
	"context"
	"errors"
	"nfxid/modules/audit/domain/hash_chain_checkpoints"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/hash_chain_checkpoints/mapper"

	"gorm.io/gorm"
)

// ByCheckpointID 根据 CheckpointID 获取 HashChainCheckpoint，实现 hash_chain_checkpoints.Get 接口
func (h *Handler) ByCheckpointID(ctx context.Context, checkpointID string) (*hash_chain_checkpoints.HashChainCheckpoint, error) {
	var m models.HashChainCheckpoint
	if err := h.db.WithContext(ctx).Where("checkpoint_id = ?", checkpointID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, hash_chain_checkpoints.ErrHashChainCheckpointNotFound
		}
		return nil, err
	}
	return mapper.HashChainCheckpointModelToDomain(&m), nil
}
