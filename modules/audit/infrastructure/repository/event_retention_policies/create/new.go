package create

import (
	"context"
	"nfxid/modules/audit/domain/event_retention_policies"
	"nfxid/modules/audit/infrastructure/repository/event_retention_policies/mapper"
)

// New 创建新的 EventRetentionPolicy，实现 event_retention_policies.Create 接口
func (h *Handler) New(ctx context.Context, erp *event_retention_policies.EventRetentionPolicy) error {
	m := mapper.EventRetentionPolicyDomainToModel(erp)
	return h.db.WithContext(ctx).Create(&m).Error
}
