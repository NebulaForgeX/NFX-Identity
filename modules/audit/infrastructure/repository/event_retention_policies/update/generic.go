package update

import (
	"context"
	"nfxid/modules/audit/domain/event_retention_policies"
	"nfxid/modules/audit/infrastructure/rdb/models"
	"nfxid/modules/audit/infrastructure/repository/event_retention_policies/mapper"
)

// Generic 通用更新 EventRetentionPolicy，实现 event_retention_policies.Update 接口
func (h *Handler) Generic(ctx context.Context, erp *event_retention_policies.EventRetentionPolicy) error {
	m := mapper.EventRetentionPolicyDomainToModel(erp)
	updates := mapper.EventRetentionPolicyModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.EventRetentionPolicy{}).
		Where("id = ?", erp.ID()).
		Updates(updates).Error
}
