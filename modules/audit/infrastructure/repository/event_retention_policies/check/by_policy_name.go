package check

import (
	"context"
	"nfxid/modules/audit/infrastructure/rdb/models"
)

// ByPolicyName 根据 PolicyName 检查 EventRetentionPolicy 是否存在，实现 event_retention_policies.Check 接口
func (h *Handler) ByPolicyName(ctx context.Context, policyName string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.EventRetentionPolicy{}).
		Where("policy_name = ?", policyName).
		Count(&count).Error
	return count > 0, err
}
