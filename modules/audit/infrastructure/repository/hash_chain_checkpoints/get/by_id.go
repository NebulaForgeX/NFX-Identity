package get

import (
	"context"
	"errors"
	"nfxid/modules/audit/domain/hash_chain_checkpoints"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/hash_chain_checkpoints/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 HashChainCheckpoint，实现 hash_chain_checkpoints.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*hash_chain_checkpoints.HashChainCheckpoint, error) {
	var m models.HashChainCheckpoint
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, hash_chain_checkpoints.ErrHashChainCheckpointNotFound
		}
		return nil, err
	}
	return mapper.HashChainCheckpointModelToDomain(&m), nil
}
