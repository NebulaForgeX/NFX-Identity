package update

import (
	"context"
	"nfxidentity/modules/audit/domain/hash_chain_checkpoints"
	"nfxidentity/modules/audit/infrastructure/rdb/models"
	"nfxidentity/modules/audit/infrastructure/repository/hash_chain_checkpoints/mapper"
)

// Generic 通用更新 HashChainCheckpoint，实现 hash_chain_checkpoints.Update 接口
func (h *Handler) Generic(ctx context.Context, hcc *hash_chain_checkpoints.HashChainCheckpoint) error {
	m := mapper.HashChainCheckpointDomainToModel(hcc)
	updates := mapper.HashChainCheckpointModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.HashChainCheckpoint{}).
		Where("id = ?", hcc.ID()).
		Updates(updates).Error
}
