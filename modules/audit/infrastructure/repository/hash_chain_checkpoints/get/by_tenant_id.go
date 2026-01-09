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

// ByTenantID 根据 TenantID 获取 HashChainCheckpoints，实现 hash_chain_checkpoints.Get 接口
func (h *Handler) ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*hash_chain_checkpoints.HashChainCheckpoint, error) {
	var ms []models.HashChainCheckpoint
	if err := h.db.WithContext(ctx).Where("tenant_id = ?", tenantID).Find(&ms).Error; err != nil {
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
