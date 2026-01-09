package create

import (
	"context"
	"nfxid/modules/audit/domain/hash_chain_checkpoints"
	"nfxid/modules/audit/infrastructure/repository/hash_chain_checkpoints/mapper"
)

// New 创建新的 HashChainCheckpoint，实现 hash_chain_checkpoints.Create 接口
func (h *Handler) New(ctx context.Context, hcc *hash_chain_checkpoints.HashChainCheckpoint) error {
	m := mapper.HashChainCheckpointDomainToModel(hcc)
	return h.db.WithContext(ctx).Create(&m).Error
}
