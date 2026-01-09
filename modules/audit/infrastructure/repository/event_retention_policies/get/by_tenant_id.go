package get

import (
	"context"
	"errors"
	"nfxid/modules/audit/domain/event_retention_policies"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/event_retention_policies/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByTenantID 根据 TenantID 获取 EventRetentionPolicies，实现 event_retention_policies.Get 接口
func (h *Handler) ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*event_retention_policies.EventRetentionPolicy, error) {
	var ms []models.EventRetentionPolicy
	if err := h.db.WithContext(ctx).Where("tenant_id = ?", tenantID).Find(&ms).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*event_retention_policies.EventRetentionPolicy{}, nil
		}
		return nil, err
	}
	
	result := make([]*event_retention_policies.EventRetentionPolicy, len(ms))
	for i := range ms {
		result[i] = mapper.EventRetentionPolicyModelToDomain(&ms[i])
	}
	return result, nil
}
