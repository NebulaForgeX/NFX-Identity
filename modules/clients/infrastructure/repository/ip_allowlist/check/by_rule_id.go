package check

import (
	"context"
	"nfxid/modules/clients/infrastructure/rdb/models"
)

// ByRuleID 根据 RuleID 检查 IPAllowlist 是否存在，实现 ip_allowlist.Check 接口
func (h *Handler) ByRuleID(ctx context.Context, ruleID string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.IpAllowlist{}).
		Where("rule_id = ?", ruleID).
		Count(&count).Error
	return count > 0, err
}
