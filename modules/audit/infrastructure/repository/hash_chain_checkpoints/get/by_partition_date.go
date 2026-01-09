package get

import (
	"context"
	"errors"
	"nfxid/modules/audit/domain/hash_chain_checkpoints"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/hash_chain_checkpoints/mapper"
	"time"

	"gorm.io/gorm"
)

// ByPartitionDate 根据 PartitionDate 获取 HashChainCheckpoints，实现 hash_chain_checkpoints.Get 接口
func (h *Handler) ByPartitionDate(ctx context.Context, partitionDate time.Time) ([]*hash_chain_checkpoints.HashChainCheckpoint, error) {
	var ms []models.HashChainCheckpoint
	if err := h.db.WithContext(ctx).Where("partition_date = ?", partitionDate).Find(&ms).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*hash_chain_checkpoints.HashChainCheckpoint{}, nil
		}
		return nil, err
	}
	
	result := make([]*hash_chain_checkpoints.HashChainCheckpoint, len(ms))
	for i := range ms {
		result[i] = mapper.HashChainCheckpointModelToDomain(&ms[i])
	}
	return result, nil
}
