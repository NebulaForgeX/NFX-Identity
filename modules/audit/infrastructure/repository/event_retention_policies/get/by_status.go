package get

import (
	"context"
	"errors"
	"nfxid/modules/audit/domain/event_retention_policies"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/event_retention_policies/mapper"

	"gorm.io/gorm"
)

// ByStatus 根据 Status 获取 EventRetentionPolicies，实现 event_retention_policies.Get 接口
func (h *Handler) ByStatus(ctx context.Context, status string) ([]*event_retention_policies.EventRetentionPolicy, error) {
	var ms []models.EventRetentionPolicy
	if err := h.db.WithContext(ctx).Where("status = ?", status).Find(&ms).Error; err != nil {
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
