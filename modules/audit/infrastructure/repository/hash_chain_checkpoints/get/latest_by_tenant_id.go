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

// LatestByTenantID 获取 TenantID 的最新 HashChainCheckpoint，实现 hash_chain_checkpoints.Get 接口
func (h *Handler) LatestByTenantID(ctx context.Context, tenantID *uuid.UUID) (*hash_chain_checkpoints.HashChainCheckpoint, error) {
	var m models.HashChainCheckpoint
	query := h.db.WithContext(ctx)
	
	if tenantID != nil {
		query = query.Where("tenant_id = ?", *tenantID)
	} else {
		query = query.Where("tenant_id IS NULL")
	}
	
	if err := query.Order("partition_date DESC, created_at DESC").First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, hash_chain_checkpoints.ErrHashChainCheckpointNotFound
		}
		return nil, err
	}
	return mapper.HashChainCheckpointModelToDomain(&m), nil
}
